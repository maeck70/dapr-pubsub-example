package main

import (
	"context"
	"fmt"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	pubsubComponentName = "orderpubsub"
	pubsubTopic         = "orders"
)

type Order_t struct {
	OrderId int    `json:"orderId"`
	Descr   string `json:"descr"`
}

func main() {
	// Create a new client for Dapr using the SDK
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Publish events using Dapr pubsub
	for i := 1; i <= 10; i++ {

		order := Order_t{
			OrderId: i,
			Descr:   "100 Led Bulbs",
		}

		err = client.PublishEvent(
			context.Background(),
			pubsubComponentName,
			pubsubTopic,
			order,
		)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Published data: %+v\n", order)

		time.Sleep(time.Second)
	}
}
