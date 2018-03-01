package main

import (
	"flag"
	"log"
	"net/http"
)

//hosting variables defined in flags
var localhost string
var localport string
var keycloakhost string
var keycloakport string
var server string
var keycloakserver string

var goTest bool // true if unit tests are running

func main() {
	flag.StringVar(&localport, "p", "3000", "Specify which port to use")
	flag.StringVar(&localhost, "host", "localhost", "Specify the name of the host")
	flag.Parse()

	server = "http://" + localhost + ":" + localport
	DatabaseInit("mcfly", "mongodb://localhost:27017/")
	router := NewRouter()
	//Stats hosting on the constant port
	log.Fatal(http.ListenAndServe(":"+localport, router))
}
