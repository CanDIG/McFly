package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var collection string

//ReadMetaData reads data from meta files
func ReadMetaData(file string) {
	b, err := ioutil.ReadFile("./uploads/" + file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	metaObject := MakeMetaObjects(s)
	//This is what determines the type of file (kind of, needs to be better [consider using file name to get type])
	if file == "meta_clinical_patient.txt" {
		collection = "patientsmeta"
	} else if file == "meta_clinical_sample.txt" {
		collection = "samplesmeta"
	} else if file == "meta_mutations_extended.txt" {
		collection = "mutationsmeta"
	} else if file == "meta_study.txt" {
		collection = "studymeta"
	}

	objects := make([]interface{}, 0)
	objects = append(objects, metaObject)
	InsertFromFile(objects, collection)
}

//ReadData reads data from cvs file
func ReadData(file string) {
	csvFile, err := os.Open("./uploads/" + file)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	objects := MakeObjects(csvData)
	//This is what determines the type of file (kind of, needs to be better [consider using file name to get type])
	if file == "data_clinical_patient.csv" {
		collection = "patients"
	} else if file == "data_clinical_sample.csv" {
		collection = "samples"
	} else if file == "data_mutations_extended.csv" {
		collection = "mutations"
	}

	InsertFromFile(objects, collection)
}

//MakeObjects creates the objects that can be stored in the database
func MakeObjects(data [][]string) []interface{} {
	data = CleanData(data)
	headers := GetHeaders(data)
	data = RemoveRow(data)
	var objects []interface{}
	for j := range data {
		m := map[string]string{}
		for i := range data[j] {
			m[strings.ToLower(headers[i])] = data[j][i]
		}
		objects = append(objects, m)
	}
	return objects
}

//MakeMetaObjects creates meta object to return
func MakeMetaObjects(s string) interface{} {
	m := map[string]string{}
	for strings.Index(s, ":") != -1 {
		i := strings.Index(s, ":")
		key := s[:i]
		s = s[i+2:]
		j := strings.Index(s, "\n")
		if j == -1 {
			value := s
			m[key] = value
			break
		}
		value := s[:j]
		s = s[j+1:]
		m[key] = value
	}
	return m
}

//GetHeaders returns the data headers
func GetHeaders(data [][]string) []string {
	return data[0]
}

//CleanData returns only the data and removes the headers
func CleanData(data [][]string) [][]string {
	for {
		if strings.Contains(data[0][0], "#") {
			data = RemoveRow(data)
		} else {
			break
		}
	}
	return data
}

//InsertFromFile inserts data records into the db
func InsertFromFile(data []interface{}, collection string) {
	for _, record := range data {
		Insert(collection, record)
	}
}

//RemoveRow removed the top row from the data
func RemoveRow(data [][]string) [][]string {
	return data[1:]
}

//MakeFileFromData exports the data in tsv file of type .txt
func MakeFileFromData(name string, data []map[string]string) {
	if len(data) == 0 {
		return
	}
	f, err := os.OpenFile("./"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringToFile := ""
	header := true
	var keys []string
	for _, result := range data {
		if header {
			keys = make([]string, 0)
			for k := range result {
				if k != "_id" {
					stringToFile += k + "\t"
					keys = append(keys, k)
				}
			}
			stringToFile += "\n"
		}
		header = false
		for _, v := range keys {
			if v != "_id" {
				stringToFile += result[v] + "\t"
			}
		}
		stringToFile += "\n"
	}
	bytes := []byte(stringToFile)

	f.Write(bytes)
	f.Close()
	return
}

//MakeSFileFromData exports the data in tsv file of type .txt
func MakeSFileFromData(name string, data []map[string]string, headers []string) {
	if len(data) == 0 {
		return
	}
	f, err := os.OpenFile("./"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	readableheaders := make([]string, 0)
	stringToFile := ""
	for _, k := range headers {
		stringToFile += k + "\t"
		readableheaders = append(readableheaders, strings.ToLower(k))
	}
	stringToFile = stringToFile[:len(stringToFile)-1]
	stringToFile += "\n"

	for l, result := range data {
		for _, v := range readableheaders {
			if v != "_id" {
				stringToFile += result[v] + "\t"
			}
		}
		stringToFile = stringToFile[:len(stringToFile)-1]
		if l != len(data)-1 {
			stringToFile += "\n"
		}
		bytes := []byte(stringToFile)
		f.Write(bytes)
		stringToFile = ""
	}
	f.Close()
	return
}

//MakeMFileFromData exports the meta data
func MakeMFileFromData(name string, data []map[string]string, headers []string) {
	if len(data) == 0 {
		return
	}
	f, err := os.OpenFile("./"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printString := ""
	for _, key := range headers {
		printString += key + ": " + data[0][key] + "\n"
	}
	printString = printString[:len(printString)-1]
	bytes := []byte(printString)
	f.Write(bytes)
	f.Close()
	return
}
