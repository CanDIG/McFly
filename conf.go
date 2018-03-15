package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

//Records is an object created by the conf.yaml file
type Records struct {
	SampleConf
	PatientConf
	MutationConf
}

//SampleConf Config Object
type SampleConf struct {
	SamplesArray []string
}

//PatientConf Config Object
type PatientConf struct {
	PatientsArray []string
}

//MutationConf is a config object
type MutationConf struct {
	MutationsArray []string
}

//GetConfSample fills the conf struct
func GetConfSample(c *SampleConf, file string) *SampleConf {
	path, _ := filepath.Abs("./" + file)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

//GetConfPatient fills the conf struct
func GetConfPatient(c *PatientConf, file string) *PatientConf {
	path, _ := filepath.Abs("./" + file)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

//GetConfMutation fills the conf struct
func GetConfMutation(c *MutationConf, file string) *MutationConf {
	path, _ := filepath.Abs("./" + file)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

//GetAllConfs grabs the arrays from the files
func GetAllConfs() *Records {
	record := Records{}
	s := &SampleConf{}
	s = GetConfSample(s, "samplesArray.yaml")
	record.SampleConf = *s
	p := &PatientConf{}
	p = GetConfPatient(p, "patientsArray.yaml")
	record.PatientConf = *p
	m := &MutationConf{}
	m = GetConfMutation(m, "mutationsArray.yaml")
	record.MutationConf = *m
	return &record
}
