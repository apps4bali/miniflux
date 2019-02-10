package gcppubsub // import "miniflux.app/integration/gcppubsub"

import (
	"fmt"
	"log"
	"context"
	"encoding/json"

	"miniflux.app/config"
	"cloud.google.com/go/pubsub"
)

// Publisher just a wrapper of pubsub Client
type Publisher struct {
	ctx context.Context
	client *pubsub.Client
	topic *pubsub.Topic
}

// NewPublisher creates new Publisher instance
func NewPublisher(config *config.Config) (publisher *Publisher) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.GcpProjectID())
	if err != nil {
		log.Fatalf("Failed top create Google PubSub client: %v", err)
	}

	topic := client.Topic(config.GcpPubsubTopic())
	return &Publisher{ctx, client, topic}
}

// Publish call pubsub client publish method
func (p *Publisher) Publish(event SyncEvent) {
	jsonEvent, err := json.Marshal(event)
	if(err != nil){
		fmt.Printf("Unable to marshal %v to JSON, %v\n", event, err)
		return
	}
	msg := &pubsub.Message{Data: []byte(jsonEvent)}
	p.topic.Publish(p.ctx, msg)
}