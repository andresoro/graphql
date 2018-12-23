package main

import (
	"fmt"

	gql "github.com/graph-gophers/graphql-go"
)

type RootResolver struct {
	db *DB
}

func (r *RootResolver) GetUser(args struct{ ID gql.ID }) {
	user, err := getUser(r.db.DB, args.ID)
	if err != nil {
		//
	}
	fmt.Print(user)
}
