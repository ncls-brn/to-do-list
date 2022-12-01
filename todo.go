package main

import (
	"time"

	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = "todo"
	port           string = ":9000"
)

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"id,omitempty"`
		Title     string        `bson:"title"`
		completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createAt"`
	}
	todo struct {
		Id        string    `json:"id"`
		Title     string    `json:"title"`
		Completed string    `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func main() {

}
