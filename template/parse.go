package template

import (
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func (template *Template) isSilentTag(token *html.Token) bool {
	if (token.Type == html.StartTagToken || token.Type == html.SelfClosingTagToken) && token.Data == "go" {
		if len(token.Attr) < 1 {
			return false
		}
		return template.SilentTags[token.Attr[0].Key]
	}
	return false
}

func (template *Template) debug(token *html.Token) {
	if template.Debug {
		template.Buffer.WriteString("\tbuffer.WriteString(`<!-- ")
		template.Buffer.WriteString(escapeBacktick(token.String()))
		template.Buffer.WriteString(" -->`)\n")
	}
}

func escapeBacktick(str string) string {
	return strings.Replace(str, "`", "`+\"`\"+`", -1)
}

const (
	// StartWrite is what starts a writing sequence in the buffer
	StartWrite       = "\tbuffer.WriteString(`"
	// CloseWrite is what ends a writing sequence in the buffer
	CloseWrite       = "`)\n"
	// StartCustomWrite is what starts a custom (ie not a strinc) sequence in the buffer (for ex. an expression)
	StartCustomWrite = "\tbuffer.WriteString("
	// ResumeWrite is usually ends a custom write and goes back to a normal writing
	ResumeWrite      = ")\n\tbuffer.WriteString(`"
)

func (template *Template) replaceExpressions(text string) string {
	var (
		result strings.Builder
		found  int
		in     bool
	)

	for {
		found = strings.Index(text, template.Expr)
		if found == -1 || len(text) < found+len(template.Expr) {
			result.WriteString(text)
			break
		}
		if !in && text[found+1:found+1+len(template.Expr)] == template.Expr {
			result.WriteString(text[:found+1])
			text = text[found+2:]
			continue
		}
		if !in {
			in = true
			result.WriteString(text[:found])
			text = text[found+1:]
			continue
		}
		coma := strings.IndexByte(text[:found], ',')
		if coma != -1 {
			format := strings.TrimSpace(text[:coma])
			escape := false
			if format[0] == '!' {
				escape = true
				format = format[1:]
			}
			expr := text[coma+1 : found]
			converted := "__fmt.Sprintf(`%" + format + "`, " + expr + ")"
			switch format {
			case "d":
				// Shortcut for int
				template.HiddenImports["strconv"] = true
				converted = `__strconv.Itoa(` + expr + `)`
			case "s":
				// Shortcut for string
				converted = expr
			default:
				template.HiddenImports["fmt"] = true
			}
			if escape {
				template.HiddenImports["html"] = true
				converted = `__html.EscapeString(` + converted + `)`
			}
			result.WriteString(CloseWrite + StartCustomWrite + converted + ResumeWrite)
		}
		text = text[found+1:]
		in = false
	}

	return result.String()
}

func (template *Template) flush() {
	if template.HTML.Len() > 0 {
		template.Buffer.WriteString(StartWrite + template.HTML.String() + CloseWrite)
	}
	template.HTML.Reset()
}

// ParseReader parse the given io.Reader
func (template *Template) ParseReader(in io.Reader) *Template {
	tokenizer := html.NewTokenizer(in)
	previousToken := html.Token{}

	for {
		if tokenizer.Next() == html.ErrorToken {
			// Returning io.EOF indicates success.
			if tokenizer.Err() != io.EOF {
				log.Println("Error", tokenizer.Err())
			}
			break
		}
		token := tokenizer.Token()

		if (token.Type == html.StartTagToken || token.Type == html.SelfClosingTagToken) && token.Data == template.Tag {
			// ----------------------------------------------------------------
			// Open or SelfClosing <go>
			// ----------------------------------------------------------------

			if len(token.Attr) == 0 {
				continue
			}
			command, exists := template.Commands[token.Attr[0].Key]
			if !exists {
				log.Println("Command", token.Attr[0].Key, "does not exists, skipping...")
				continue
			}
			template.debug(&token)
			template.flush()
			command(template, &token)

		} else if token.Type == html.EndTagToken && token.Data == template.Tag {
			// ----------------------------------------------------------------
			// Close </go>
			// ----------------------------------------------------------------

			template.debug(&token)
			template.flush()
			template.Buffer.WriteString("}\n")

		} else if token.Type == html.TextToken {
			// ----------------------------------------------------------------
			// Text
			// ----------------------------------------------------------------

			if template.isSilentTag(&previousToken) {
				data := strings.TrimSpace(token.Data)
				if data == "" {
					continue
				}
			}
			text := escapeBacktick(token.Data)
			if template.Trim {
				text = strings.TrimSpace(text)
			}
			text = template.replaceExpressions(text)
			template.HTML.WriteString(text)
		} else {
			// ----------------------------------------------------------------
			// Other tags, looking for $expression$
			// ----------------------------------------------------------------

			tag := escapeBacktick(token.String())
			tag = template.replaceExpressions(tag)
			template.HTML.WriteString(tag)
		}
		previousToken = token
	}
	template.flush()
	return template
}

// ParseFile parse the given file
func (template *Template) ParseFile(file string) *Template {
	in, err := os.Open(file)
	if err != nil {
		log.Println("Unable to open", file)
		log.Fatal(err)
	}
	defer in.Close()
	return template.ParseReader(in)
}

// ParseString parse the given string
func (template *Template) ParseString(data string) *Template {
	return template.ParseReader(strings.NewReader(data))
}
