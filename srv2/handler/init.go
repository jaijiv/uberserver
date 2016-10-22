package handler

import (
	"log"

	"github.com/jaijiv/uberserver/srv2/subscriber"
	"github.com/micro/go-micro"

	example "github.com/jaijiv/uberserver/srv2/proto/example"
)

func RunMain() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.srv2"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(Example))

	// Register Struct as Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber("topic.go.micro.srv.srv2", new(subscriber.Example)),
	)

	// Register Function as Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber("topic.go.micro.srv.srv2", subscriber.Handler),
	)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
