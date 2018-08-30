package template

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

// Command represents the default signature of any commands
type Command func(template *Template, token *html.Token)

// Template represents a full template with its configuration and parsed buffer
type Template struct {
	Package       string
	Imports       map[string]bool
	HiddenImports map[string]bool
	FuncName      string
	FuncArgs      string
	Buffer        *bytes.Buffer
	Commands      map[string]Command
	SilentTags    map[string]bool
	Trim          bool
	Debug         bool
	HTML          strings.Builder
	Expr	      string
	Tag       	  string
}
