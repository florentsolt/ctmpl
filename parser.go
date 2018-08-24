package gotmpl

import (
	// "bytes"
	"log"
	"os"
	"io"
	"golang.org/x/net/html"
)

type Template struct {
	Imports []string
	FuncName string
	FuncArgs string  
}

func ParseFile(filename string) {
	fd, err := os.Open(filename)
	if err != nil {

	}
	tokenizer := html.NewTokenizer(fd)
	for {
		if tokenizer.Next() == html.ErrorToken {
			// Returning io.EOF indicates success.
			if tokenizer.Err() != io.EOF {
				log.Println("Error", tokenizer.Err())
			}
			return
		}
		log.Printf("Token %#v\n", tokenizer.Token())
	}	
}