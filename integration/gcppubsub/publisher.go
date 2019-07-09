package gcppubsub // import "miniflux.app/integration/gcppubsub"

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"miniflux.app/config"
	"miniflux.app/timer"
)

// Publisher just a wrapper of pubsub Client
type Publisher struct {
	ctx    context.Context
	client *pubsub.Client
	topic  *pubsub.Topic
}

// NewPublisher creates new Publisher instance
func NewPublisher(config *config.Config) (publisher *Publisher) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.GcpProjectID())
	if err != nil {
		log.Fatalf("[gcppubsub:NewPublisher] Failed to create Google PubSub client: %v\n", err)
	}

	topic := client.Topic(config.GcpPubsubTopic())
	return &Publisher{ctx, client, topic}
}

// PublishEvent publish an event to PubSub
func (p *Publisher) PublishEvent(event SyncEvent) {
	jsonEvent, err := json.Marshal(event)
	if err != nil {
		log.Printf("[Publisher:PublishEvent] Unable to marshal %v to JSON, %v\n", event, err)
		return
	}
	msg := &pubsub.Message{Data: []byte(jsonEvent)}

	// TODO: Context should not inside a Struct
	_, err = p.topic.Publish(p.ctx, msg).Get(p.ctx)
	if err != nil {
		log.Printf("[Publisher:PublishEvent] Publishing to topic failed, %v", err)
		return
	}
	timer.ExecutionTime(time.Now(), fmt.Sprintf("[Publisher:PublishEvent] Publishing %v", event))
}
