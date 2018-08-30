package main

import (
	"bytes"
	"os"
)

func main() {
	buffer := new(bytes.Buffer)
	Index(buffer)
	buffer.WriteTo(os.Stdout)
}
