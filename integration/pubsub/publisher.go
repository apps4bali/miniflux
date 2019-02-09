package pubsub // import "miniflux.app/integration/pubsub"

import (
	"fmt"
	"log"
	"context"
	"encoding/json"

	"miniflux.app/config"
	gpubsub "cloud.google.com/go/pubsub"
)

// Publisher just a wrapper of pubsub Client
type Publisher struct {
	ctx context.Context
	client *gpubsub.Client
	topic *gpubsub.Topic
}

// NewPublisher creates new Publisher instance
func NewPublisher(config *config.Config) (publisher *Publisher) {
	ctx := context.Background()
	client, err := gpubsub.NewClient(ctx, config.GcpProjectID())
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
	msg := &gpubsub.Message{Data: []byte(jsonEvent)}
	p.topic.Publish(p.ctx, msg)
}