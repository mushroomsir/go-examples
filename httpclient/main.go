package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	"github.com/siddontang/go/bson"
)

func main() {
	var cc interface{}
	cc = []string{}
	c, ok := cc.([]interface{})
	log.Println(c, ok)

	var projectID bson.ObjectId
	fmt.Println(projectID.Hex())
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Hour)
	}))
	defer svr.Close()
	fmt.Println("making request")
	//	http.Get(svr.URL)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	_, err := client.Get(svr.URL)
	if e, ok := err.(*url.Error); ok && e.Timeout() {
		fmt.Println("url.Error:", e)
	}

	fmt.Println("finished request")
}
