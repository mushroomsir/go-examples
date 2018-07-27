package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "code.teambition.com/soa/go-lib/pkg/health"
)

// Result ...
type Result struct {
	r   *http.Response
	err error
}

func main() {
	time.Sleep(time.Hour)
	//process()
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	resultChan := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.xxx.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		resultChan <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		er := <-resultChan
		fmt.Println("Timeout:", er.err)
	case res := <-resultChan:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}
