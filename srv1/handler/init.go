package handler

import (
	"log"

	"github.com/jaijiv/uberserver/srv1/subscriber"
	"github.com/micro/go-micro"

	example "github.com/jaijiv/uberserver/srv1/proto/example"
)

func RunMain() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.srv1"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(Example))

	// Register Struct as Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber("topic.go.micro.srv.srv1", new(subscriber.Example)),
	)

	// Register Function as Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber("topic.go.micro.srv.srv1", subscriber.Handler),
	)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
