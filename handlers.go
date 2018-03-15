package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
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
	newFileName := header.Filename[0:len(header.Filename)-4] + ".csv"
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
	ReadData(newFileName)
	return
})

//patientJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var patientJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	objects := GetAll("patients")
	if err := json.NewEncoder(w).Encode(objects); err != nil {
		panic(err)
	}
	//MakePatientFileFromData("results.txt", objects)
	MakeSFileFromData("results.txt", objects, Arrays.PatientConf.PatientsArray)
})

//sampleJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var sampleJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	objects := GetAll("samples")
	if err := json.NewEncoder(w).Encode(objects); err != nil {
		panic(err)
	}
	//MakeSampleFileFromData("results.txt", objects)
	MakeSFileFromData("results.txt", objects, Arrays.SampleConf.SamplesArray)
})

//mutationJSONHandler returns a comma seperated list of generic JSON objects stored in objects array
var mutationJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	objects := GetAll("mutations")
	// if err := json.NewEncoder(w).Encode(objects); err != nil {
	// 	panic(err)
	// }
	w.Write([]byte("results.txt created"))
	//MakeMutationFileFromData("results.txt", objects)
	MakeSFileFromData("results.txt", objects, Arrays.MutationConf.MutationsArray)
})
