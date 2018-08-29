
		package main
		import (
			"bytes"
			"os"
		)

		func main() {
			buffer := new(bytes.Buffer)
			Basic(
				"mytitle", 
				"myclass",
				true,
				[]string{"item1", "item2", "item3"},
				map[string]string{"k":"v"},
				42,
				42.42424242,
				buffer,
			)
			buffer.WriteTo(os.Stdout)
		}