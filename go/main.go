package main

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"main/user/grpc"
	"math/rand"
	"net"
)

func main() {
	app := fiber.New()
	grcpServer := grpc.NewServer()
	var server Server
	user.RegisterUserServer(grcpServer, server)
	listen, err := net.Listen("tcp", "localhost:3000")
	if err != nil{
		log.Fatal("Could not listen to localhost:3000 %v", err)
	}
	go grcpServer.Serve(listen)
	log.Println("GRCP Server Started")
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	app.Listen(":4000")
	log.Println("Fiber Server Started")

}

type Server struct {}



func (s Server) Search(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {

	data := user.UserResponse{Found: true, Id: int32(rand.Float32()), Username: request.Username}
	return &data, nil
}
