package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"

	"github.com/dgraph-io/badger"
	"github.com/graph-gophers/graphql-go"
)

type User struct {
	Id   graphql.ID
	Name string
}

func addUser(db *badger.DB, u User) error {

	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, u)
	b, err := ioutil.ReadAll(&buf)
	if err != nil {
		return err
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(u.Id), b)
		if err != nil {
			return err
		}
		_ = txn.Commit()
		return nil
	})
	return nil
}
