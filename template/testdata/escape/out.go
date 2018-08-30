package main

import (
	"bytes"
	"os"
)

func main() {
	buffer := new(bytes.Buffer)
	Escape("<script></script>", buffer)
	buffer.WriteTo(os.Stdout)
}
