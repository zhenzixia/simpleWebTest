package main

import (
	"net/http"
	"log"
	"simpleWebTest/pkg/resource"
)

func main() {
	service := &resource.RecordResource{}
	service.Initialize("13.59.145.88:8125", "test01")
	service.Register()
	log.Println("Starting Http Server!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
