package user

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserreq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return 
	}

	res, err := h.Service.CreateUser(c,&u)
	if err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(200, res)
}


func (h *Handler) LoginUser(c *gin.Context){
	var user LoginReq
	if err :=c.ShouldBindJSON(&user); err!= nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	u,err := h.Service.Loginuser(c,&user)
	if err != nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	c.SetCookie("jwt",u.AccessToken,3600,"/","localhost",false,true)

	res := &LoginRes{
		Username: u.Username,
		Id: u.Id,
	}

	c.JSON(200,res)
}

func (h *Handler) Logout(c *gin.Context){
	c.SetCookie("jwt","",-1,"","",false,true)
	c.JSON(200,gin.H{"message":"logout success"})
}