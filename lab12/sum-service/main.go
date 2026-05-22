package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	totalSum := 0

	_, err = nc.Subscribe("pipeline.squared", func(m *nats.Msg) {
		var msg Message
		if err := json.Unmarshal(m.Data, &msg); err != nil {
			log.Printf("Unmarshalling error: %v", err)
			return
		}
		if msg.Value == -1 {
			log.Println("Summing done. Service stopped")
			log.Printf("Total sum: %d ", totalSum)
			err = nc.Flush()
			if err != nil {
				log.Fatalf("Flush error: %v", err)
			}
			os.Exit(0)
		}

		totalSum += msg.Value
	})
	if err != nil {
		log.Fatalf("Subscribe error: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
