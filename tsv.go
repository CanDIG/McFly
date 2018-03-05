package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
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
func MakeObjects(data [][]string) []Record {
	data = CleanData(data)
	headers := GetHeaders(data)
	data = RemoveRow(data)
	var objects []Record
	for j := range data {
		m := map[string]string{}
		for i := range data[j] {
			m[headers[i]] = data[j][i]
		}

		record := TranslateData(m)
		objects = append(objects, record)
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
func InsertFromFile(data []Record) {
	for _, record := range data {
		Insert("data", record)
	}
}

//RemoveRow removed the top row from the data
func RemoveRow(data [][]string) [][]string {
	return data[1:]
}

//MakeFileFromData exports the data in tsv file of type .txt
func MakeFileFromData(name string, data []Record) {
	if len(data) == 0 {
		return
	}
	f, err := os.OpenFile("./"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringToFile := ""
	r := Record{}
	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		tag, _ := reflect.ValueOf(r).Type().FieldByName(v.Type().Field(i).Name)
		stringToFile += string(tag.Tag.Get("json")) + "\t"
	}
	stringToFile = stringToFile[:len(stringToFile)-1]
	stringToFile += "\n"

	for _, k := range data {
		h := reflect.ValueOf(k)
		for i := 0; i < v.NumField(); i++ {
			value := fmt.Sprintf("%v", h.Field(i))
			stringToFile += value + "\t"
		}
		stringToFile = stringToFile[:len(stringToFile)-1]
		stringToFile += "\n"
	}
	stringToFile = stringToFile[:len(stringToFile)-1]
	bytes := []byte(stringToFile)

	f.Write(bytes)
	f.Close()
	return
}

//TranslateData is a function that translate into the recordobejct
func TranslateData(m map[string]string) Record {
	r := Record{}
	tag, _ := reflect.ValueOf(r).Type().FieldByName("Age")
	r.Age, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("AngiolymphaticInvasion")
	r.AngiolymphaticInvasion = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("ClinicalAssessment")
	r.ClinicalAssessment = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("DfsMonths")
	r.DfsMonths = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("DfsStatus")
	r.DfsStatus = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Ethnicity")
	r.Ethnicity = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("LostToFollowup")
	r.LostToFollowup = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("LymphNodesExamined")
	r.LymphNodesExamined = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MethodOfInitialPathDiagnosis")
	r.MethodOfInitialPathDiagnosis = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MethodOfInitialPathDiagnosisOther")
	r.MethodOfInitialPathDiagnosisOther = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("NewTumorEventAfterInitialTreatment")
	r.NewTumorEventAfterInitialTreatment = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("NewTumorEventOtherSite")
	r.NewTumorEventOtherSite = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("OsMonths")
	r.OsMonths = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("OsStatus")
	r.OsStatus = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("PathologicDistantSpread")
	r.PathologicDistantSpread = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("PathologicSpreadIncidentalProstateCancer")
	r.PathologicSpreadIncidentalProstateCancer = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("PatientID")
	r.PatientID = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("PrimaryTumorPathologicSpread")
	r.PrimaryTumorPathologicSpread = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("ProspectiveCollection")
	r.ProspectiveCollection = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Race")
	r.Race = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Sex")
	r.Sex = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("SmokingHistory")
	r.SmokingHistory = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("SmokingPackYears")
	r.SmokingPackYears, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("SmokingYearStarted")
	r.SmokingYearStarted = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("SmokingYearStopped")
	r.SmokingYearStopped = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Subtype")
	r.Subtype = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("YearOfDiagnosis")
	r.YearOfDiagnosis = m[string(tag.Tag.Get("json"))]

	return r
}
