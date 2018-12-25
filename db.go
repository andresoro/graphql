package main

import (
	"github.com/dgraph-io/badger"
)

// DB
type DB struct {
	*badger.DB
}

// NewDB returns new db
func NewDB() (*DB, error) {

	opts := badger.DefaultOptions
	opts.Dir = "./db/badger"
	opts.ValueDir = "./db/badger"

	db, err := badger.Open(opts)
	if err != nil {
		return &DB{}, err
	}

	return &DB{db}, nil
}
