package main

import (
	"fmt"
	"log"
    "github.com/siddontang/go-mysql/canal"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "mysql-server:3306"
	cfg.User = "root"
	cfg.Password = "root"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "mydb"
	cfg.Dump.Tables = []string{"mytable"}
	
	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Println(err)
	}
	
	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})
	
	// Start canal
	c.Run()

}

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnTableChanged(schema string, table string) error {
	url := "mongodb-server"
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	collection := session.DB("mymongodb").C("mymongocollection")
	log.Println("Removing all")
	collection.RemoveAll(bson.M{})
	return nil
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	url := "mongodb-server"
	session, err := mgo.Dial(url)
	if err != nil {
		log.Println(err)
	}
	log.Println(session)
	collection := session.DB("mymongodb").C("mymongocollection")
	log.Println(collection)

	log.Println(e.Table.Columns)
	log.Println("Action", e.Action)

	for _, row := range e.Rows {
		if e.Action == "insert" || e.Action == "update" {
			var value string
			switch v := row[1].(type) {
			case string:
				value = v
			case []byte:
				value = string(v)
			}
			fmt.Println(row[1])
			data := map[string]interface{}{
				"_id": row[0],
				"value": value,
			}
			log.Println("upserting...")
			_, err := collection.Upsert(bson.M{"_id": row[0]}, data)
			if err != nil {
				log.Println(err)
			}
		} else if e.Action == "delete" {
			err := collection.Remove(bson.M{"_id": row[0]})
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

