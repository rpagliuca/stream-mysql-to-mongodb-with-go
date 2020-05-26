package main

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"log"
	"fmt"
)

func main() {
	url := "mongodb-server"
	session, err := mgo.Dial(url)
	if err != nil {
		log.Println(err)
	}
	log.Println(session)
	collection := session.DB("mymongodb").C("mymongocollection")
	rows := []interface{}{}
	err = collection.Find(bson.M{}).All(&rows)
	if err != nil {
		log.Println(err)
	}
	for _, row := range rows {
		fmt.Println(row)
	}
}
