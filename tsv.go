package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

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
	InsertFromFile(objects)
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
			m[headers[i]] = data[j][i]
		}
		objects = append(objects, m)
	}
	return objects
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
func InsertFromFile(data []interface{}) {
	for _, record := range data {
		Insert("data", record)
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
