

# package template
`import "github.com/florentsolt/gotmpl/template"`





## Index 
- [Constants]L(#constants)
- [Variables]L(#variables)
- [Functions]L(#functions)
  - [func CmdElse(template *Template, token *html.Token)]L(#func-CmdElse)
  - [func CmdFor(template *Template, token *html.Token)]L(#func-CmdFor)
  - [func CmdFunc(template *Template, token *html.Token)]L(#func-CmdFunc)
  - [func CmdIf(template *Template, token *html.Token)]L(#func-CmdIf)
  - [func CmdImport(template *Template, token *html.Token)]L(#func-CmdImport)
  - [func CmdInclude(template *Template, token *html.Token)]L(#func-CmdInclude)
  - [func CmdPackage(template *Template, token *html.Token)]L(#func-CmdPackage)


- [Types]L(#types)
  - [type Command]L(#type-Command)


  - [type Template]L(#type-Template)
    - [func New() *Template]L(#func-New)

    - [func (template *Template) ParseFile(file string) *Template]L(#func-TemplateParseFile)
    - [func (template *Template) ParseReader(in io.Reader) *Template]L(#func-TemplateParseReader)
    - [func (template *Template) ParseString(data string) *Template]L(#func-TemplateParseString)
    - [func (template *Template) Save(file string)]L(#func-TemplateSave)







## Constants


``` go
const (
    // StartWrite is what starts a writing sequence in the buffer
    StartWrite = "\tbuffer.WriteString(`"
    // CloseWrite is what ends a writing sequence in the buffer
    CloseWrite = "`)\n"
    // StartCustomWrite is what starts a custom (ie not a strinc) sequence in the buffer (for ex. an expression)
    StartCustomWrite = "\tbuffer.WriteString("
    // ResumeWrite is usually ends a custom write and goes back to a normal writing
    ResumeWrite = ")\n\tbuffer.WriteString(`"
)
```

## Variables


``` go
var DefaultCommands = map[string]LCommand{
    "package": CmdPackage,
    "import":  CmdImport,
    "func":    CmdFunc,
    "include": CmdInclude,
    "if":      CmdIf,
    "else":    CmdElse,
    "for":     CmdFor,
}
```
DefaultCommands represents the default set of commands (as a tag attribute)
available by default

## Functions



### func CmdElse
``` go
func CmdElse(template *Template, token *html.Token)
```
CmdElse represents the tag <go else />



### func CmdFor
``` go
func CmdFor(template *Template, token *html.Token)
```
CmdFor represents the tag <go for="expression"> Note that it's not auto closing,
you will need to close it with </go>



### func CmdFunc
``` go
func CmdFunc(template *Template, token *html.Token)
```
CmdFunc represents the tag <go func="Name" args="a string, b int,..." />



### func CmdIf
``` go
func CmdIf(template *Template, token *html.Token)
```
CmdIf represents the tag <go if="expression"> Note that it's not auto closing,
you will need to close it with </go>



### func CmdImport
``` go
func CmdImport(template *Template, token *html.Token)
```
CmdImport represents the tag <go import="name" />



### func CmdInclude
``` go
func CmdInclude(template *Template, token *html.Token)
```
CmdInclude represents the tag <go include="Name" args="a, b,..." />



### func CmdPackage
``` go
func CmdPackage(template *Template, token *html.Token)
```
CmdPackage represents the tag <go package="name" />

## Types



### type Command
``` go
type Command func(template *Template, token *html.Token)```
Command represents the default signature of any commands



### type Template
``` go
type Template struct {
    Package       string
    Imports       map[string]Lbool
    HiddenImports map[string]Lbool
    FuncName      string
    FuncArgs      string
    Buffer        *bytes.Buffer
    Commands      map[string]LCommand
    SilentTags    map[string]Lbool
    Trim          bool
    Debug         bool
    HTML          strings.Builder
    Expr          string
    Tag           string
}
```
Template represents a full template with its configuration and parsed buffer




#### func New
``` go
func New() *Template
```
New returns a fresh template ready use with default values




#### func Template.ParseFile
``` go
func (template *Template) ParseFile(file string) *Template
```
ParseFile parse the given file





#### func Template.ParseReader
``` go
func (template *Template) ParseReader(in io.Reader) *Template
```
ParseReader parse the given io.Reader





#### func Template.ParseString
``` go
func (template *Template) ParseString(data string) *Template
```
ParseString parse the given string





#### func Template.Save
``` go
func (template *Template) Save(file string)
```
Save write the generated go file in the given filename




