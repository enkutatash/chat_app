package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		hub: hub,
	}
}

type CreateRoomReq struct{
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var roomReq CreateRoomReq
	if err:= c.ShouldBindJSON(&roomReq); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	h.hub.Room[roomReq.Id] = &Room{
		Id: roomReq.Id,
		Name: roomReq.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(200,roomReq)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request)bool{
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context){
	conn,err := upgrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	roomId := c.Param("roomid")
	clientId := c.Query("clientid")
	username := c.Query("username")

	cl := &Client{
		Conn: conn,
		Message: make(chan *Message,10),
		Id: clientId,
		RoomId: roomId,
		Username: username,
	}
	n:=&Message{
		Content: "new user joined",
		RoomId: roomId,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- n

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}

type RoomRes struct{
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
func (h *Handler) GetRooms(c *gin.Context){
 rooms := make([]RoomRes,0)
 for _,r := range h.hub.Room{
	rooms = append(rooms,RoomRes{
		Id: r.Id,
		Name: r.Name,
	})
 }

 c.JSON(200,rooms)
}

type ClientRes struct{
 Id string `json:"id" db:"id"`
 Username string `json:"username" db:"username"`
}


func (h *Handler) GetClients(c *gin.Context){
 var clients []ClientRes
 roomId := c.Param("roomid")
 if _,ok := h.hub.Room[roomId];!ok{
	clients = make([]ClientRes,0)
	c.JSON(200,clients)
 }
for _,c:= range(h.hub.Room[roomId].Clients){
	clients = append(clients,ClientRes{
		Id: c.Id,
		Username: c.Username,
	})
 }
 c.JSON(200,clients)
}