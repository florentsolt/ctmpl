package template

import (
	"log"

	"golang.org/x/net/html"
)

// DefaultCommands represents the default set of commands (as a tag attribute) available by default
var DefaultCommands = map[string]Command{
	"package": CmdPackage,
	"import":  CmdImport,
	"func":    CmdFunc,
	"include": CmdInclude,
	"if":      CmdIf,
	"else":    CmdElse,
	"for":     CmdFor,
}

// CmdPackage represents the tag <go package="name" />
func CmdPackage(template *Template, token *html.Token) {
	template.Package = token.Attr[0].Val
}

// CmdImport represents the tag <go import="name" />
func CmdImport(template *Template, token *html.Token) {
	template.Imports[token.Attr[0].Val] = true
}

// CmdFunc represents the tag <go func="Name" args="a string, b int,..." />
func CmdFunc(template *Template, token *html.Token) {
	// FIXME auto capitalize ?
	template.FuncName = token.Attr[0].Val
	if len(token.Attr) > 1 && token.Attr[1].Key == "args" {
		template.FuncArgs = token.Attr[1].Val
	}
}

// CmdInclude represents the tag <go include="Name" args="a, b,..." />
func CmdInclude(template *Template, token *html.Token) {
	if len(token.Attr) > 1 && token.Attr[1].Key == "args" {
		template.Buffer.WriteString("\t " + token.Attr[0].Val + "(" + token.Attr[1].Val + ", buffer)\n")
	} else {
		template.Buffer.WriteString("\t " + token.Attr[0].Val + "(buffer)\n")
	}
}

// CmdIf represents the tag <go if="expression">
// Note that it's not auto closing, you will need to close it with </go>
func CmdIf(template *Template, token *html.Token) {
	if token.Type == html.SelfClosingTagToken {
		log.Fatal("<go if> should not auto close")
	}
	template.Buffer.WriteString("\tif " + token.Attr[0].Val + " {\n")
}

// CmdElse represents the tag <go else />
func CmdElse(template *Template, token *html.Token) {
	if token.Type != html.SelfClosingTagToken {
		log.Fatal("<go else> should auto close")
	}
	template.Buffer.WriteString("\t} else {\n")
}

// CmdFor represents the tag <go for="expression">
// Note that it's not auto closing, you will need to close it with </go>
func CmdFor(template *Template, token *html.Token) {
	if token.Type == html.SelfClosingTagToken {
		log.Fatal("<go for> should not auto close")
	}
	template.Buffer.WriteString("\tfor " + token.Attr[0].Val + " {\n")
}
