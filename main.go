package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "http port")
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Printf("server start at %d",*port)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err!=nil{
		log.Printf("error happed: %v",err)
	}
}
