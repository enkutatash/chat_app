package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string,args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}


func (r *repository) CreateUser(ctx context.Context,user *User)(*User,error){
	var userId int
	query := "INSERT INTO userstable(username,email,password) VALUES($1,$2,$3) RETURNING id"
	err :=r.db.QueryRowContext(ctx,query,user.Username,user.Email,user.Password).Scan(&userId)
	if err != nil{
		return &User{},nil
	}
	user.Id = int64(userId)
	return user,nil
	
}

func (r *repository) GetUserByEmail(ctx context.Context,email string)(*User,error){
	var user User;
	query := "SELECT id,username, email,password FROM userstable WHERE email = $1"
	err := r.db.QueryRowContext(ctx,query,email).Scan(&user.Id,&user.Username,&user.Email,&user.Password)
	if err != nil{
		return &User{},nil
	}
	return &user,nil
}