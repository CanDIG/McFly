package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

var databaseName string
var connectionString string

var session *mgo.Session

//DatabaseInit creates a connection to the database
func DatabaseInit(dbName, connectionstring string) {
	databaseName = dbName
	connectionString = connectionstring + dbName

	session, err = mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
}

func setCollection(collection string) *mgo.Collection {
	return session.DB(databaseName).C(collection)
}

//Insert allows users to add generic objects to a collection in the database
func Insert(collection string, object interface{}) bool {
	c := setCollection(collection)
	err := c.Insert(object)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//GetAll returns an array of all objects in a collection
func GetAll(collection string) []Record {
	c := setCollection(collection)
	var list []Record
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//RemoveAll will empty a collection
func RemoveAll(collection string) bool {
	c := setCollection(collection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
