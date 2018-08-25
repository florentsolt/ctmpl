
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
