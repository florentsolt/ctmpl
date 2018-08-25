package template

import (
	"bytes"

	"golang.org/x/net/html"
)

type Command func(template *Template, token *html.Token)

type Template struct {
	Package  string 
	Imports  map[string]bool
	FuncName string
	FuncArgs string
	Buffer *bytes.Buffer
	Commands map[string]Command
}