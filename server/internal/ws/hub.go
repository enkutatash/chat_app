package ws

type Room struct{
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Clients map[string]*Client `json:"clients" db:"clients"`
}

type Hub struct {
	Room map[string]*Room `json:"room" db:"room"`
	Register chan *Client
	Unregister chan *Client
	Broadcast chan *Message
}

func NewHub() *Hub{
	return &Hub{
		Room: make(map[string]*Room),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast: make(chan *Message,5),
	}
}

func (h *Hub) Run(){
	for{
		select{
		case cl:= <-h.Register:
			if _,ok := h.Room[cl.RoomId];ok{
				r:= h.Room[cl.RoomId]
				if _,ok := r.Clients[cl.Id];!ok{
					r.Clients[cl.Id] = cl
				}
			}
		case cl:= <-h.Unregister:
			if _,ok := h.Room[cl.RoomId];ok{
				if _,ok := h.Room[cl.RoomId].Clients[cl.Id];ok{
					if len(h.Room[cl.RoomId].Clients) != 0{
						h.Broadcast <- &Message{
							Content: cl.Username + " has left the room",
							RoomId: cl.RoomId,
							Username: cl.Username,
						}
					}
					delete(h.Room[cl.RoomId].Clients,cl.Id)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _,ok := h.Room[m.RoomId];ok{
				for _,cl := range h.Room[m.RoomId].Clients{
					cl.Message <- m
				}
			}
		}
	}
}