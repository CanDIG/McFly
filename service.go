package main

//MakeAFile creates a file as a service
func MakeAFile(collection string, name string, headers []string, isData bool, head string) {
	objects := GetAll(collection)
	if isData {
		MakeSFileFromData(name, objects, headers, head)
	} else {
		MakeMFileFromData(name, objects, headers)
	}
	return
}
