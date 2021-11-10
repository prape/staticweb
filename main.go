package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	currentPath,err := os.Getwd()
	if err!=nil{
		currentPath="."
	}
	port := flag.Int("port", 8080, "http port")
	webpath := flag.String("path",currentPath,"static web path")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*webpath)))
	log.Printf("server %s on port %d",*webpath,*port)
	err = http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err!=nil{
		log.Printf("error happed: %v",err)
	}
}
