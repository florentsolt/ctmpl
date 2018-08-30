package main

import (
	"fmt"
	"github.com/florentsolt/gotmpl/template"
	"io/ioutil"
	"log"
	"strings"
	"flag"
)

func main() {
	trim := flag.Bool("trim", false, "Enable trim")
	pkg := flag.String("package", "", "Specifiy package name")
	expr := flag.String("expr", "$", "Specifiy epxression token")
	tag := flag.String("tag", "go", "Specifiy tag name")
	flag.Parse()

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			fmt.Println(file.Name())
			template := template.New()
			template.Trim = *trim
			template.Expr = *expr
			template.Tag = *tag
			template.ParseFile(file.Name())
			if *pkg != "" {
				template.Package = *pkg
			}
			template.Save(file.Name() + ".go")
		}
	}
}
