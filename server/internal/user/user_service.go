package user

import (
	"context"
	"fmt"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	SECRETKEY = "secret"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserreq) (*CreateUserRes, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	//  password hash
	hashPassword, err := util.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	}
	r, err := s.Repository.CreateUser(c, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		Id:       strconv.Itoa(int(r.Id)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil
}

type JWTClaim struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}



func (s *service) Loginuser(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.Repository.GetUserByEmail(c, req.Email)

	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	err = util.CheckPassword( req.Password,user.Password)

	if err != nil {
		return nil, fmt.Errorf("password don't match")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		Id:       strconv.Itoa(int(user.Id)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.Id)),
			ExpiresAt:  jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(SECRETKEY))

	if err != nil {
		return nil, err
	}

	res := &LoginRes{
		AccessToken: ss,
		Id:          strconv.Itoa(int(user.Id)),
		Username:    user.Username,
	}

	return res, nil

}
