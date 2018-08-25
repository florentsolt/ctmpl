package main

import (
	"testing"
	"os"

	"github.com/florentsolt/gotmpl/template"
)

var code = `
	package main
	import (
		"bytes"
		"os"
		"github.com/florentsolt/gotmpl/template/testdata"
	)

	func main() {
		buffer := new(bytes.Buffer)
		testdata.Test1(
			"title", 
			"class",
			true,
			[]string{"item1", "item2", "item3"},
			map[string]string{"k":"v"},
			buffer,
		)
		buffer.WriteTo(os.Stdout)
	}
` 

func TestParseFile(t *testing.T) {
	template := template.New()
	template.Parse("../template/testdata/test1.html")
	template.Package = "testdata" 
	template.Save("../template/testdata/test1.html.go")

	file := "testdata/test1.go"
	out, err := os.Create(file)
	if err != nil {
		t.Error("Unable to open", file)
		t.Error(err)
	}
	defer out.Close()
	out.WriteString(code)
	
	t.Fail()
}
