package main

import (
	"github.com/dgraph-io/badger"
)

type Resolver struct {
	db *badger.DB
}
