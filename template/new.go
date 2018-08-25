package template

import (
	"bytes"
)

func New() *Template {
	return &Template{
		Imports: map[string]bool{"bytes":true}, 
		Buffer: new(bytes.Buffer),
		Commands: DefaultCommands,
	}
}