package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// make a sample HTTP GET request
	res, err := http.Get( "https://jsonplaceholder.typicode.com/users/1" )

	// check for response error
	if err != nil {
		log.Fatal( err )
	}

	// read all response body
	data, _ := ioutil.ReadAll( res.Body )

	// close response body
	res.Body.Close()

	// print request `Content-Type` header
	requestContentType := res.Request.Header.Get( "Content-Type" )
	contentType := res.Header.Get( "Content-Type" )
	fmt.Println( "Request content-type: {}; content-type: {}", requestContentType, contentType )

	// print `data` as a string
	fmt.Printf( "%s\n", data )
}