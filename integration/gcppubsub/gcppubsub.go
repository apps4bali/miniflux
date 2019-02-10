package gcppubsub // import "miniflux.app/integration/gcppubsub"

// Constants related to SyncEvent
const (
	EntityTypeCategory string = "CATEGORY"
	EntityTypeFeed string = "FEED"
	EntityTypeEntry string = "ENTRY"

	EntityOpWrite string = "WRITE"
	EntityOpDelete string = "DELETE"
)

// SyncEvent model
type SyncEvent struct {
	EntityType string `json:"entity_type"`
	EntityID int64 `json:"entity_id"`
	EntityOp string `json:"entity_op"`
}

// NewCategoryEvent returns `SyncEvent` with type `EntityTypeCategory`
func NewCategoryEvent(categoryID int64, op string) SyncEvent {
	return SyncEvent{EntityTypeCategory, categoryID, op}
}

// NewFeedEvent returns `SyncEvent` with type `EntityTypeFeed`
func NewFeedEvent(feedID int64, op string) SyncEvent {
	return SyncEvent{EntityTypeFeed, feedID, op}
}

// NewEntryEvent returns `SyncEvent` with type `EntityTypeEntry`
func NewEntryEvent(entryID int64, op string) SyncEvent {
	return SyncEvent{EntityTypeEntry, entryID, op}
}