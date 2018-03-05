package main

import "log"

//GetAllPatients returns an array of all objects in a collection
func GetAllPatients(collection string) []PatientRecord {
	c := setCollection(collection)
	var list []PatientRecord
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//GetAllSamples returns an array of all objects in a collection
func GetAllSamples(collection string) []SampleRecord {
	c := setCollection(collection)
	var list []SampleRecord
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//GetAllMutations returns an array of all objects in a collection
func GetAllMutations(collection string) []MutationRecord {
	c := setCollection(collection)
	var list []MutationRecord
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}
