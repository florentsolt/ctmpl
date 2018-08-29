package template

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

type Command func(template *Template, token *html.Token)

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
	Html          strings.Builder
}

// EnableTrim enables trim mode, all useless spaces will be removed
// Needs to be called before any Parse methods
func (template *Template) EnableTrim() *Template {
	template.Trim = true
	return template
}

// EnableDebug enables debug mode, a comment will precede all <go> tags substitutions
// Needs to be called before any Parse methods
func (template *Template) EnableDebug() *Template {
	template.Debug = true
	return template
}
