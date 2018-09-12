package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//log.Println(bson.NewObjectId())
	go func() {
		defer recoverPanic()
		go func() { panic(fmt.Errorf("This unhandled error will be handled")) }()
		time.Sleep(10 * time.Second)
	}()
	time.Sleep(10 * time.Second)
}

func recoverPanic() {
	if rec := recover(); rec != nil {
		log.Println(rec)
		err := rec.(error)
		fmt.Printf("Unhandled error: %v\n", err.Error())
		fmt.Fprintf(os.Stderr, "Program quit unexpectedly; please check your logs\n")
		os.Exit(1)
	}
}
