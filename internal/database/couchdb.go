package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-kivik/couchdb/v3"
	"github.com/go-kivik/kivik/v3"
)

var Client *kivik.Client

func InitCouchDB() {
    client, err := kivik.New("couch", "http://admin:password@localhost:5984/")
    if err != nil {
        log.Fatal("Failed to connect to CouchDB: ", err)
    }
	fmt.Println("Database connected successfully")
    Client = client
}

func GetDB(databaseName string) *kivik.DB {
    return Client.DB(context.TODO(), databaseName)
}


