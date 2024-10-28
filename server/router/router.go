package router

import (
	"server/internal/user"
	"server/internal/ws"
	"github.com/gin-gonic/gin"
)

var(
	r *gin.Engine
)

func InitRouter(userHandler *user.Handler,wsHandler *ws.Handler){
	r = gin.Default()
	r.POST("/signup",userHandler.CreateUser)
	r.POST("/login",userHandler.LoginUser)
	r.GET("/logout",userHandler.Logout)
	r.POST("/ws/createroom",wsHandler.CreateRoom)
	r.GET("/ws/joinroom/:roomid",wsHandler.JoinRoom)
	r.GET("/ws/getrooms",wsHandler.GetRooms)
	r.GET("ws/getclients/:roomid",wsHandler.GetClients)
}

func Start(addr string) error{
	return r.Run(addr)
}