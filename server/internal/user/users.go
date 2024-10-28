package user

import "context"

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context,email string)(*User,error)

}

type Service interface{
  CreateUser(ctx context.Context,req *CreateUserreq)(*CreateUserRes, error)
  Loginuser(ctx context.Context,req *LoginReq) (*LoginRes, error)
}

type CreateUserreq struct{
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type CreateUserRes struct{
	Id       string  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
type LoginReq struct{
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginRes struct{
	AccessToken string
	Id       string  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}