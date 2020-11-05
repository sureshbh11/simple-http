package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// handle `/` route
	http.HandleFunc( "/", func( res http.ResponseWriter, req *http.Request ) {
			fmt.Fprint(res, "Hello World!")
		},
	)

	log.Output(1, "Starting to listen...Ready")

	// run server on port "9000"
	log.Fatal( http.ListenAndServe( ":9000", nil ) )

}