package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

//Global vairable definitions
var err error

//Index returns when the main page is called and returns HTML indicating the availale paths
var indexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.New("").Delims("[[", "]]").ParseFiles("./views/index.html")
	//temp, _ := tpl.Parse(tplfile)
	tpl.ExecuteTemplate(w, "index.html", nil)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return
})

var uploadHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	var newFileName string
	if !strings.Contains(header.Filename, "meta") {
		newFileName = header.Filename[0:len(header.Filename)-4] + ".csv"
	} else {
		newFileName = header.Filename
	}
	if _, err := os.Stat("./uploads/" + newFileName); !os.IsNotExist(err) {
		w.WriteHeader(409)
		return
	}
	f, err := os.OpenFile("./uploads/"+newFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
	}
	if !strings.Contains(header.Filename, "meta") {
		ReadData(newFileName)
	} else {
		ReadMetaData(newFileName)
	}
	return
})

//patientJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var patientJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("patients file created"))

	MakeAFile("patients", "data_clinical_patient.txt", Arrays.PatientConf.PatientsArray, true)
})

//sampleJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var sampleJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("samples file created"))

	MakeAFile("samples", "data_clinical_sample.txt", Arrays.SampleConf.SamplesArray, true)
})

//mutationJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var mutationJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("mutations file created"))

	MakeAFile("mutations", "data_mutations_extended.txt", Arrays.MutationConf.MutationsArray, true)
})

//patientJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var patientmetaJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("patient metadata file created"))

	MakeAFile("patientsmeta", "meta_clinical_patient.txt", Arrays.PatientMetaConf.PatientsMetaArray, false)
})

//sampleJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var samplemetaJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("sample metadata file created"))

	MakeAFile("samplesmeta", "meta_clinical_sample.txt", Arrays.SampleMetaConf.SamplesMetaArray, false)
})

//mutationJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var mutationmetaJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("mutation metadata file created"))

	MakeAFile("mutationsmeta", "meta_mutations_extended.txt", Arrays.MutationMetaConf.MutationsMetaArray, false)
})

//mutationJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var studymetaJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("study metadata file created"))

	MakeAFile("studymeta", "meta_study.txt", Arrays.StudyConf.StudyArray, false)
})

var makeAllHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Creating all files"))

	MakeAFile("patients", "data_clinical_patient.txt", Arrays.PatientConf.PatientsArray, true)
	MakeAFile("samples", "data_clinical_sample.txt", Arrays.SampleConf.SamplesArray, true)
	MakeAFile("mutations", "data_mutations_extended.txt", Arrays.MutationConf.MutationsArray, true)
	MakeAFile("patientsmeta", "meta_clinical_patient.txt", Arrays.PatientMetaConf.PatientsMetaArray, false)
	MakeAFile("samplesmeta", "meta_clinical_sample.txt", Arrays.SampleMetaConf.SamplesMetaArray, false)
	MakeAFile("mutationsmeta", "meta_mutations_extended.txt", Arrays.MutationMetaConf.MutationsMetaArray, false)
	MakeAFile("studymeta", "meta_study.txt", Arrays.StudyConf.StudyArray, false)
})
