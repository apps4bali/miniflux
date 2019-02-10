// Copyright 2017 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package storage // import "miniflux.app/storage"

import (
	"database/sql"
	
	"miniflux.app/integration/gcppubsub"
)

// Storage handles all operations related to the database.
type Storage struct {
	db *sql.DB
	pub *gcppubsub.Publisher
}

// NewStorage returns a new Storage.
func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

// AddPubsubPublisher sets the pub to the Storage instance
func (s *Storage) AddPubsubPublisher(pub *gcppubsub.Publisher) {
	s.pub = pub
}