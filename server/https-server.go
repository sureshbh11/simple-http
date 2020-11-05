package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
)

type Job struct {
	Id int
	Name string
	Status string
}

func main() {

	// generate a `Certificate` struct
	cert, _ := tls.LoadX509KeyPair( "localhost.crt", "localhost.key" )

	// create a custom server
	s := &http.Server{
		Addr: ":9002",
		Handler: nil, // use `http.DefaultServeMux`
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{ cert },
		},
	}

	// handle `/listJobs` route
	http.HandleFunc( "/listJobs", func( res http.ResponseWriter, req *http.Request ) {

		jobs := []Job{

			Job{
				100,
				"Job-100",
				"r",
			},
			Job{
				101,
				"Job-101",
				"i",
			},
		}

		// encode `john` as JSON
		jobJSON, err := json.MarshalIndent(jobs, "", " " )

		// print JSON string
		fmt.Fprint(res, string(jobJSON), err )

		//		fmt.Fprint(res, "Mock Slurm Job Scheduler Jobs")
	},
	)

//	id := 100
	// handle `/job` route
//	http.HandleFunc( fmt.Sprint("/job?id=%d", id), func( res http.ResponseWriter, req *http.Request ) {
	http.HandleFunc( "/job", func( res http.ResponseWriter, req *http.Request ) {

		id := req.URL.Query().Get("id")
		ids, err := strconv.Atoi(id)
		job := Job{
			ids,
			fmt.Sprint("Job-",id),
			"r",
		}

		// encode `john` as JSON
		jobJSON, err := json.MarshalIndent(job, "", " " )

		// print JSON string
		fmt.Fprint(res, string(jobJSON), err )

//		fmt.Fprint(res, "Mock Slurm Job Scheduler Jobs")
	},
	)

	log.Output(1, "Mock Slurm Job Scheduler Starting to listening on 9002...Ready")

	// run server on port "9002"
	log.Fatal( s.ListenAndServeTLS( "localhost.crt", "localhost.key" ) )

}