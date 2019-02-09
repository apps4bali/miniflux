package pubsub // import "miniflux.app/integration/pubsub"

// Constants related to SyncEvent
const (
	EntityTypeCategory = "CATEGORY"
	EntityTypeFeed = "FEED"
	EntityTypeEntry = "ENTRY"

	EntityOpWrite = "WRITE"
	EntityOpDelete = "DELETE"
)

// SyncEvent ...
type SyncEvent struct {
	EntityID int `json:"entity_id"`
	EntityType string `json:"entity_type"`
	EntityOp string `json:"entity_op"`
}