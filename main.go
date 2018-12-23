package main

import (
	"io/ioutil"
	"log"
	"net/http"

	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {

	//load db
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := getSchema("./schema.graphql")

	schema := gql.MustParseSchema(s, &RootResolver{db: db})
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
