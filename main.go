package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgraph-io/badger"

	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {

	//db
	opts := badger.DefaultOptions
	opts.Dir = "./db/badger"
	opts.ValueDir = "./db/badger"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := getSchema("./schema.graphql")

	schema := gql.MustParseSchema(s, &Resolver{db: db})
	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getSchema(path string) string {
	f, err := ioutil.ReadFile("./schema.graphql")
	if err != nil {
		log.Fatal(err)
	}

	return string(f)
}
