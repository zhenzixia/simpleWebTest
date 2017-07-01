package main

import (
	"net/http"
	"log"
	"simpleWebTest/pkg/resource"
)

func main() {
	service := &resource.ProductResource{}
	service.Initialize("http://13.59.145.88:8125", "test01")
	service.Register()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}*/
