package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Job struct {
	Id int
	Name string
	Status string
}

func main() {

	// create a new HTTP client
	client := &http.Client {
		Timeout: 2 * time.Minute,
	}

	// send an GET request from `client`
//	res, err := client.Get( "https://localhost:9002/listJobs" )
	res, err := client.Get( "https://localhost:9002/job?id=100" )

	// check for response error
	if err != nil {

		// get `url.Error` struct pointer from `err` interface
		urlErr := err.( *url.Error )

		// check if error occurred due to timeout
		if urlErr.Timeout() {
			fmt.Println( "Error occurred due to a timeout." );
		}

		// log error and exit
		log.Println( "Error:", err )
	} else {
		fmt.Println( "Success: status-code", res.StatusCode );
	}

	// read all response body
	data, _ := ioutil.ReadAll( res.Body )

	// close response body
	res.Body.Close()

	// print request `Content-Type` header
	requestContentType := res.Request.Header.Get( "Content-Type" )
	contentType := res.Header.Get( "Content-Type" )
	fmt.Println( "Request content-type: {}; content-type: {}", requestContentType, contentType )

	size := len(data)
	data = data[:size-5]
	var job Job
	// unmarshal `data`
	err = json.Unmarshal( data, &job );
	if err != nil {
		// log error and exit
		log.Println("Error:", err)
	}

	// print `data` as a string
	fmt.Printf( "%s\n", data )
}