package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/websocket/v2"
	"github.com/slipe-fun/skid-backend/internal/app"
	"github.com/slipe-fun/skid-backend/internal/config"
	"github.com/slipe-fun/skid-backend/internal/repository"
	"github.com/slipe-fun/skid-backend/internal/service"
	"github.com/slipe-fun/skid-backend/internal/transport/http"
)

func main() {
	cfg := config.LoadConfig("configs/config.yaml")

	db := repository.InitDB(cfg)
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	chatRepo := repository.NewChatRepo(db)

	jwtSvc := service.NewJWTService(cfg.JWT.Secret)
	tokenSvc := service.NewTokenService(jwtSvc)

	authApp := app.NewAuthApp(userRepo, jwtSvc)
	userApp := app.NewUserApp(userRepo, jwtSvc, tokenSvc)
	chatApp := app.NewChatApp(chatRepo, tokenSvc)

	authHandler := http.NewAuthHandler(authApp)
	userHandler := http.NewUserHandler(userApp)
	chatHandler := http.NewChatHandler(chatApp, userApp)

	fiberApp := fiber.New()

	fiberApp.Post("/auth/login", authHandler.Login)
	fiberApp.Post("/auth/register", authHandler.Register)

	fiberApp.Get("/user/me", userHandler.GetUser)
	fiberApp.Get("/user/:id", userHandler.GetUserById)

	fiberApp.Post("/chat/create", chatHandler.CreateChat)
	fiberApp.Get("/chats", chatHandler.GetChatsByUserId)
	fiberApp.Get("/chat/:id", chatHandler.GetChatById)

	// fiberApp.Get("/ws", websocket.New(func(c *websocket.Conn) {
	// 	defer c.Close()
	// 	for {
	// 		mt, msg, err := c.ReadMessage()
	// 		if err != nil {
	// 			break
	// 		}
	// 		c.WriteMessage(mt, msg)
	// 	}
	// }))

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)))
}
