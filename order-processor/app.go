package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

var sub = &common.Subscription{
	PubsubName: "orderpubsub",
	Topic:      "orders",
	Route:      "/orders",
}

func main() {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "6005"
	}

	// Create the new server on appPort and add a topic listener
	s := daprd.NewService(":" + appPort)
	err := s.AddTopicEventHandler(sub, eventHandler)
	if err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	// Start the server
	err = s.Start()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listening: %v", err)
	}
}

type Order_t struct {
	OrderId int    `json:"orderId"`
	Descr   string `json:"descr"`
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	// Use Dapr to Unmarshal the data into the Order_t struct
	var data Order_t
	err = e.Struct(&data)
	if err != nil {
		fmt.Println("Error Struct data:", err)
		return false, err
	}

	// Show what we have received
	fmt.Printf("COntentType: %s\n", e.DataContentType)
	fmt.Printf("Raw data: %s\n", e.RawData)
	fmt.Printf("OrderId: %d\n", data.OrderId)
	fmt.Printf("Descr: %s\n", data.Descr)

	return false, nil
}
