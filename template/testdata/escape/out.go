
	package main
	import (
		"bytes"
		"os"
	)

	func main() {
		buffer := new(bytes.Buffer)
		Escape(buffer)
		buffer.WriteTo(os.Stdout)
	}