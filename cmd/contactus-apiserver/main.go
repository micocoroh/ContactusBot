package main

import (
	"log"
	"net"

	"github.com/micocoroh/contactus/api"
	"github.com/micocoroh/contactus/internal/service"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	chatbotService := service.NewChatbotService()

	api.RegisterChatbotServiceServer(server, chatbotService)
	server.Serve(listenPort)
}
