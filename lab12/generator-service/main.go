package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

type Message struct {
	Value int `json:"value"`
}

func main() {
	natsURL := os.Getenv("NATS_URL")

	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer nc.Close()

	time.Sleep(3 * time.Second)

	for i := 1; i <= 100; i++ {
		msg := Message{Value: i}
		data, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Marshaling error : %v", err)
			continue
		}

		err = nc.Publish("pipeline.numbers", data)
		if err != nil {
			log.Printf("Publication error: %v", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	err = nc.Publish("pipeline.numbers", []byte(`{"value":-1}`))
	if err != nil {
		log.Fatalf("Publishing error: %v", err)
	}

	fmt.Println("Generation done. Service stopped")
}
