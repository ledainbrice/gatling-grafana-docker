package main

import "sample-service2/app/services"
import "fmt"

type Container struct {}

func NewContainer() *Container {
    return &Container{}
}

func (c *Container) Check() {
    fmt.Printf("Container is set.")
}

func (c *Container) CreateEvent() services.Event {
    return services.NewEvent(c.CreateGreeter())
}

func (c *Container) CreateGreeter() services.Greeter {
    return services.NewGreeter(c.CreateMessage())
}

func (c *Container) CreateMessage() services.Message {
    return services.NewMessage()
}