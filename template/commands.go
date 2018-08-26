package template

import (
	"log"

	"golang.org/x/net/html"
)

var DefaultCommands = map[string]Command{
	"package": CmdPackage,
	"import":  CmdImport,
	"func":    CmdFunc,
	"json":    CmdJson,
	"if":      CmdIf,
	"for":     CmdFor,
}

func CmdPackage(template *Template, token *html.Token) {
	template.Package = token.Attr[0].Val
}

func CmdImport(template *Template, token *html.Token) {
	template.Imports[token.Attr[0].Val] = true
}

func CmdFunc(template *Template, token *html.Token) {
	// FIXME auto capitalize ?
	template.FuncName = token.Attr[0].Val
	if len(token.Attr) > 1 && token.Attr[1].Key == "args" {
		template.FuncArgs = token.Attr[1].Val
	}
}

func CmdString(template *Template, token *html.Token) {
	template.Buffer.WriteString("\tbuffer.WriteString(")
	template.Buffer.WriteString(token.Attr[0].Val)
	template.Buffer.WriteString(")\n")
}

func CmdInt(template *Template, token *html.Token) {
	template.Imports["strconv"] = true
	template.Buffer.WriteString("\tbuffer.WriteString(strconv.Itoa(")
	template.Buffer.WriteString(token.Attr[0].Val)
	template.Buffer.WriteString("))\n")
}

func CmdJson(template *Template, token *html.Token) {
	template.Imports["fmt"] = true
	template.Buffer.WriteString("\tbuffer.WriteString(`<pre><code>`)\n")
	template.Buffer.WriteString("\tbuffer.WriteString(fmt.Sprintf(`%#v`, " + token.Attr[0].Val + "))\n")
	template.Buffer.WriteString("\tbuffer.WriteString(`</code></pre>`)\n")
}

func CmdIf(template *Template, token *html.Token) {
	if token.Type == html.SelfClosingTagToken {
		log.Fatal("<go if> should not auto close")
	}
	template.Buffer.WriteString("\tif " + token.Attr[0].Val + " {\n")
}

func CmdFor(template *Template, token *html.Token) {
	if token.Type == html.SelfClosingTagToken {
		log.Fatal("<go for> should not auto close")
	}
	template.Buffer.WriteString("\tfor " + token.Attr[0].Val + " {\n")
}
