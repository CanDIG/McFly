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

	reader.Comma = '\t' // Use tab-delimited instead of comma

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	csvData = CleanData(csvData)

	//This is what determines the type of file (kind of, needs to be better [consider using file name to get type])
	if csvData[0][1] == "AGE" {
		objects := MakePatients(csvData)
		InsertPatientFromFile(objects)
	} else if csvData[0][1] == "SAMPLE_ID" {
		objects := MakeSamples(csvData)
		InsertSampleFromFile(objects)
	} else if csvData[0][0] == "Hugo_Symbol" {
		objects := MakeMutations(csvData)
		InsertMutationFromFile(objects)
	}

	return
}

//MakePatients creates the objects that can be stored in the database
func MakePatients(data [][]string) []PatientRecord {
	headers := GetHeaders(data)
	data = RemoveRow(data)
	var objects []PatientRecord
	for j := range data {
		m := map[string]string{}
		for i := range data[j] {
			m[headers[i]] = data[j][i]
		}

		record := TranslatePatientData(m)
		objects = append(objects, record)
	}
	return objects
}

//MakeSamples creates the objects that can be stored in the database
func MakeSamples(data [][]string) []SampleRecord {
	headers := GetHeaders(data)
	data = RemoveRow(data)
	var objects []SampleRecord
	for j := range data {
		m := map[string]string{}
		for i := range data[j] {
			m[headers[i]] = data[j][i]
		}

		record := TranslateSampleData(m)
		objects = append(objects, record)
	}
	return objects
}

//MakeMutations creates the objects that can be stored in the database
func MakeMutations(data [][]string) []MutationRecord {
	headers := GetHeaders(data)
	data = RemoveRow(data)
	var objects []MutationRecord
	for j := range data {
		m := map[string]string{}
		for i := range data[j] {
			m[headers[i]] = data[j][i]
		}

		record := TranslateMutationData(m)
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

//InsertPatientFromFile inserts data records into the db
func InsertPatientFromFile(data []PatientRecord) {
	for _, record := range data {
		Insert("patients", record)
	}
}

//InsertSampleFromFile inserts data records into the db
func InsertSampleFromFile(data []SampleRecord) {
	for _, record := range data {
		Insert("samples", record)
	}
}

//InsertMutationFromFile inserts data records into the db
func InsertMutationFromFile(data []MutationRecord) {
	for _, record := range data {
		Insert("mutations", record)
	}
}

//RemoveRow removed the top row from the data
func RemoveRow(data [][]string) [][]string {
	return data[1:]
}

//MakePatientFileFromData exports the data in tsv file of type .txt
func MakePatientFileFromData(name string, data []PatientRecord) {
	if len(data) == 0 {
		return
	}
	f, err := os.OpenFile("./"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringToFile := ""
	r := PatientRecord{}
	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		tag, _ := reflect.ValueOf(r).Type().FieldByName(v.Type().Field(i).Name)
		stringToFile += string(tag.Tag.Get("json")) + "\t"
	}
	stringToFile = stringToFile[:len(stringToFile)-1]
	stringToFile += "\n"

	for l, k := range data {
		h := reflect.ValueOf(k)
		for i := 0; i < v.NumField(); i++ {
			value := fmt.Sprintf("%v", h.Field(i))
			stringToFile += value + "\t"
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
	f.Close()
	return
}

//MakeSampleFileFromData exports the data in tsv file of type .txt
func MakeSampleFileFromData(name string, data []SampleRecord) {
	if len(data) == 0 {
		return
	}
	f, err := os.OpenFile("./"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringToFile := ""
	r := SampleRecord{}
	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		tag, _ := reflect.ValueOf(r).Type().FieldByName(v.Type().Field(i).Name)
		stringToFile += string(tag.Tag.Get("json")) + "\t"
	}
	stringToFile = stringToFile[:len(stringToFile)-1]
	stringToFile += "\n"

	for l, k := range data {
		h := reflect.ValueOf(k)
		for i := 0; i < v.NumField(); i++ {
			value := fmt.Sprintf("%v", h.Field(i))
			stringToFile += value + "\t"
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

//MakeMutationFileFromData exports the data in tsv file of type .txt
func MakeMutationFileFromData(name string, data []MutationRecord) {
	if len(data) == 0 {
		return
	}
	f, err := os.OpenFile("./"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringToFile := ""
	r := MutationRecord{}
	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		tag, _ := reflect.ValueOf(r).Type().FieldByName(v.Type().Field(i).Name)
		stringToFile += string(tag.Tag.Get("json")) + "\t"
	}
	stringToFile = stringToFile[:len(stringToFile)-1]
	stringToFile += "\n"

	for l, k := range data {
		h := reflect.ValueOf(k)
		for i := 0; i < v.NumField(); i++ {
			value := fmt.Sprintf("%v", h.Field(i))
			stringToFile += value + "\t"
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

//TranslatePatientData is a function that translate into the recordobejct
func TranslatePatientData(m map[string]string) PatientRecord {
	r := PatientRecord{}
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

//TranslateSampleData is a function that translate into the recordobejct
func TranslateSampleData(m map[string]string) SampleRecord {
	r := SampleRecord{}
	tag, _ := reflect.ValueOf(r).Type().FieldByName("BatchNumber")
	r.BatchNumber, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("CancerTypeDetailed")
	r.CancerTypeDetailed, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("CompleteTCGAID")
	r.CompleteTCGAID, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("ExtracapsularSpread")
	r.ExtracapsularSpread, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Grade")
	r.Grade, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("IfExtracapsularSpreadPresent")
	r.IfExtracapsularSpreadPresent, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("IncidentalProstateCancer")
	r.IncidentalProstateCancer, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MetastaticSite")
	r.MetastaticSite, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MetastaticSiteOther")
	r.MetastaticSiteOther, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("NewTumorEventSite")
	r.NewTumorEventSite, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("NewTumorEventType")
	r.NewTumorEventType, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("NodesEx")
	r.NodesEx, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("OncotreeCode")
	r.OncotreeCode, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("PatientID")
	r.PatientID, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("PrimarySite")
	r.PrimarySite, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("PriorTumor")
	r.PriorTumor, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("RegionalNodesPathologicSpread")
	r.RegionalNodesPathologicSpread, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("SampleID")
	r.SampleID, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("TStage")
	r.TStage, _ = m[string(tag.Tag.Get("json"))]

	return r
}

//TranslateMutationData is a function that translate into the recordobejct
func TranslateMutationData(m map[string]string) MutationRecord {
	r := MutationRecord{}
	tag, _ := reflect.ValueOf(r).Type().FieldByName("BAMFile")
	r.BAMFile, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Center")
	r.Center, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Chromosome")
	r.Chromosome, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Consequence")
	r.Consequence, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("DBSNPRS")
	r.DBSNPRS, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("DBSNPValStatus")
	r.DBSNPValStatus, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("EndPosition")
	r.EndPosition, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("EntrezGeneID")
	r.EntrezGeneID, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("HGVSc")
	r.HGVSc, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("HGVSp")
	r.HGVSp, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("HGVSpShort")
	r.HGVSpShort, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("HugoSymbol")
	r.HugoSymbol, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MatchedNormSampleBarcode")
	r.MatchedNormSampleBarcode, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MatchNormSeqAllele1")
	r.MatchNormSeqAllele1, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MatchNormSeqAllele2")
	r.MatchNormSeqAllele2, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MatchNormValidationAllele1")
	r.MatchNormValidationAllele1, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MatchNormValidationAllele2")
	r.MatchNormValidationAllele2, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("MutationStatus")
	r.MutationStatus, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("NAltCount")
	r.NAltCount, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("NCBIBuild")
	r.NCBIBuild, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("NRefCount")
	r.NRefCount, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("ReferenceAllele")
	r.ReferenceAllele, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Score")
	r.Score, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Sequencer")
	r.Sequencer, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("SequenceSource")
	r.SequenceSource, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("SequencingPhase")
	r.SequencingPhase, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("StartPosition")
	r.StartPosition, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("Strand")
	r.Strand, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("TAltCount")
	r.TAltCount, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("TRefCount")
	r.TRefCount, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("TumorSampleBarcode")
	r.TumorSampleBarcode, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("TumorSeqAllele1")
	r.TumorSeqAllele1, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("TumorSeqAllele2")
	r.TumorSeqAllele2, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("TumorValidationAllele1")
	r.TumorValidationAllele1, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("TumorValidationAllele2")
	r.TumorValidationAllele2, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("ValidationMethod")
	r.ValidationMethod, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("ValidationStatus")
	r.ValidationStatus, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("VariantClassification")
	r.VariantClassification, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("VariantType")
	r.VariantType, _ = m[string(tag.Tag.Get("json"))]
	tag, _ = reflect.ValueOf(r).Type().FieldByName("VerificationStatus")
	r.VerificationStatus, _ = m[string(tag.Tag.Get("json"))]

	return r
}
