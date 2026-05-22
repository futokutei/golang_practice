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

	_, err = nc.Subscribe("pipeline.numbers", func(m *nats.Msg) {
		var msg Message
		if err := json.Unmarshal(m.Data, &msg); err != nil {
			log.Printf("Unmarshalling error: %v", err)
			return
		}
		if msg.Value == -1 {
			log.Println("Filtering done. Service stopped")
			err := nc.Publish("pipeline.even", m.Data)
			if err != nil {
				log.Fatalf("End error: %v", err)
			}
			err = nc.Flush()
			if err != nil {
				log.Fatalf("Flush error: %v", err)
			}
			os.Exit(0)
		}

		if msg.Value%2 == 0 {
			err := nc.Publish("pipeline.even", m.Data)
			if err != nil {
				log.Fatalf("Publishing error: %v", err)
				return
			}
		}
	})
	if err != nil {
		log.Fatalf("Subscribe error: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
