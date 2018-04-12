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
	SampleMetaConf
	PatientConf
	PatientMetaConf
	MutationConf
	MutationMetaConf
	StudyConf
}

//SampleConf Config Object
type SampleConf struct {
	SamplesArray []string
	Headers      string
}

//PatientConf Config Object
type PatientConf struct {
	PatientsArray []string
	Headers       string
}

//MutationConf is a config object
type MutationConf struct {
	MutationsArray []string
	Headers        string
}

//SampleMetaConf Config Object
type SampleMetaConf struct {
	SamplesMetaArray []string
}

//PatientMetaConf Config Object
type PatientMetaConf struct {
	PatientsMetaArray []string
}

//MutationMetaConf is a config object
type MutationMetaConf struct {
	MutationsMetaArray []string
}

//StudyConf is a config object
type StudyConf struct {
	StudyArray []string
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
	c.Headers = "#McFly\n"
	return c
}

//GetConfMetaSample fills the conf struct
func GetConfMetaSample(c *SampleMetaConf, file string) *SampleMetaConf {
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

//GetConfMetaPatient fills the conf struct
func GetConfMetaPatient(c *PatientMetaConf, file string) *PatientMetaConf {
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

//GetConfMetaMutation fills the conf struct
func GetConfMetaMutation(c *MutationMetaConf, file string) *MutationMetaConf {
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

//GetConfStudy fills the conf struct
func GetConfStudy(c *StudyConf, file string) *StudyConf {
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

	sm := &SampleMetaConf{}
	sm = GetConfMetaSample(sm, "meta_samples.yaml")
	record.SampleMetaConf = *sm
	pm := &PatientMetaConf{}
	pm = GetConfMetaPatient(pm, "meta_patients.yaml")
	record.PatientMetaConf = *pm
	mm := &MutationMetaConf{}
	mm = GetConfMetaMutation(mm, "meta_mutations.yaml")
	record.MutationMetaConf = *mm

	ms := &StudyConf{}
	ms = GetConfStudy(ms, "meta_study.yaml")
	record.StudyConf = *ms
	return &record
}
