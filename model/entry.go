// Copyright 2017 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package model // import "miniflux.app/model"

import (
	"fmt"
	"time"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Entry statuses
const (
	EntryStatusUnread       = "unread"
	EntryStatusRead         = "read"
	EntryStatusRemoved      = "removed"
	DefaultSortingOrder     = "published_at"
	DefaultSortingDirection = "asc"
)

// Entry represents a feed item in the system.
type Entry struct {
	ID          int64         `json:"id"`
	UserID      int64         `json:"user_id"`
	FeedID      int64         `json:"feed_id"`
	Status      string        `json:"status"`
	Hash        string        `json:"hash"`
	Title       string        `json:"title"`
	URL         string        `json:"url"`
	CommentsURL string        `json:"comments_url"`
	Date        time.Time     `json:"published_at"`
	Content     string        `json:"content"`
	Author      string        `json:"author"`
	Starred     bool          `json:"starred"`
	Enclosures  EnclosureList `json:"enclosures,omitempty"`
	Feed        *Feed         `json:"feed,omitempty"`
	Category    *Category     `json:"category,omitempty"`
}

// Entries represents a list of entries.
type Entries []*Entry


// GetEnclosuresFromContent returns EnclosureList by finding specific tag on entry content
// and convert them into Enclosure object.
// But this enclosure doesn't actually exists on DB, for client side use only.
func (e Entry) GetEnclosuresFromContent() EnclosureList {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(e.Content))
	if err != nil {
		return nil
	}
	spans := doc.Find("span[data-miniflux-enclosure]")
	list := make(EnclosureList, 0)

	if spans.Length() > 0 {
		spans.Each(func(i int, span *goquery.Selection) {
			url, _ := span.Attr("data-miniflux-enclosure")
			tokens := strings.Split(url, ".")
			extension := tokens[len(tokens)-1]
			list = append(list, &Enclosure{
					URL: url,
					MimeType: "image/" + strings.ToLower(extension),
					Size: 0,
				})
		})
	}

	return list
}

// ValidateEntryStatus makes sure the entry status is valid.
func ValidateEntryStatus(status string) error {
	switch status {
	case EntryStatusRead, EntryStatusUnread, EntryStatusRemoved:
		return nil
	}

	return fmt.Errorf(`Invalid entry status, valid status values are: "%s", "%s" and "%s"`, EntryStatusRead, EntryStatusUnread, EntryStatusRemoved)
}

// ValidateEntryOrder makes sure the sorting order is valid.
func ValidateEntryOrder(order string) error {
	switch order {
	case "id", "status", "published_at", "category_title", "category_id":
		return nil
	}

	return fmt.Errorf(`Invalid entry order, valid order values are: "id", "status", "published_at", "category_title", "category_id"`)
}

// ValidateDirection makes sure the sorting direction is valid.
func ValidateDirection(direction string) error {
	switch direction {
	case "asc", "desc":
		return nil
	}

	return fmt.Errorf(`Invalid direction, valid direction values are: "asc" or "desc"`)
}

// ValidateRange makes sure the offset/limit values are valid.
func ValidateRange(offset, limit int) error {
	if offset < 0 {
		return fmt.Errorf(`Offset value should be >= 0`)
	}

	if limit < 0 {
		return fmt.Errorf(`Limit value should be >= 0`)
	}

	return nil
}

// OppositeDirection returns the opposite sorting direction.
func OppositeDirection(direction string) string {
	if direction == "asc" {
		return "desc"
	}

	return "asc"
}
