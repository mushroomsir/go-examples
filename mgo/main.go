package main

import (
	"fmt"
	"log"

	"github.com/siddontang/go/bson"
	mgo "gopkg.in/mgo.v2"
)

// Person ...
type Person struct {
	Age   int    `bson:"age,omitempty"`
	Name  string `bson:"name,omitempty"`
	Phone string `bson:"phone,omitempty"`
}

func main() {
	log.Println(bson.NewObjectId())
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Collection People
	c := session.DB("test123").C("people")

	query := bson.M{
		"name": "abc",
	}
	update := bson.M{
		"$set": bson.M{"age": 1},
		"$setOnInsert": bson.M{
			"name":  "abcc",
			"phone": "186",
		},
	}
	change := mgo.Change{
		Update:    update,
		Upsert:    true,
		ReturnNew: true,
	}
	result := &Person{}
	_, err = c.Find(query).Apply(change, result)
	fmt.Println(err)
}
