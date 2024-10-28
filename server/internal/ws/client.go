package ws

import (
	"log"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	Id       string `json:"id" db:"id"`
	RoomId   string `json:"room_id" db:"room_id"`
	Username string `json:"username" db:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomId   string `json:"room_id"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage(){
	defer func() {
		c.Conn.Close()
	}()

	for {
		message,ok := <-c.Message
		if !ok{
			return 
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(h *Hub){
	defer func(){
		h.Unregister <- c
		c.Conn.Close()
	}()

	for{
		_,m,err :=c.Conn.ReadMessage()
		if err != nil{
			if websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
				log.Printf("error: %v",err)
			}
			break
		}
		msg := &Message{
			Content: string(m),
			RoomId: c.RoomId,
			Username: c.Username,
		}

		h.Broadcast <- msg
	}
}