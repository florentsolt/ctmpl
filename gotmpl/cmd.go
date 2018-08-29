package main

import (
	"fmt"
	"github.com/florentsolt/gotmpl/template"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			fmt.Println(file.Name())
			template := template.New()
			template.ParseFile(file.Name())
			template.Save(file.Name() + ".go")
		}
	}
}
