package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

/*
  @Description:
*/

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// file create
	f,err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	multi := io.MultiWriter(f,os.Stdout)

	io.Copy(multi,r.Body)

	if err := r.Body; err != nil {
		log.Fatalln(err)
	}

}
