package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
	"server/internal/ws"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("couldn't init db", err)
	}
	userRepo := user.NewRepository(dbConn.GetDB())
	userSev := user.NewService(userRepo)
	userHandler := user.NewHandler(userSev)

	hub:= ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler,wsHandler)
	router.Start("0.0.0.0:8080")

}