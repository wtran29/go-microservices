package main

import (
	"context"
	"log"
	"time"

	"github.com/wtran29/go-microservices/logger/data"
)

// RPCServer is the type for our RPC server. Methods that uses this as a receiver are available over RPC as long
// as long as they're exported.
type RPCServer struct{}

// RPCPayload is the type for data we received from RPC
type RPCPayload struct {
	Name string
	Data string
}

// LogInfo wriets the payload to mongo
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}
	*resp = "Processed payload via RPC:" + payload.Name
	return nil
}
