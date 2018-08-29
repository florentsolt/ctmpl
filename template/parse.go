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
	CloseWrite       = "`)\n"
	StartWrite       = "\tbuffer.WriteString(`"
	StartCustomWrite = "\tbuffer.WriteString("
	ResumeWrite      = ")\n\tbuffer.WriteString(`"
)

func (template *Template) replaceVariables(text string) string {
	var (
		result strings.Builder
		found  int
		in     bool
	)

	for {
		found = strings.IndexByte(text, '$')
		if found == -1 || len(text) < found+1 {
			result.WriteString(text)
			break
		}
		if !in && text[found+1] == '$' {
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
			converted := "fmt.Sprintf(`%" + format + "`, " + expr + ")"
			switch format {
			case "d":
				template.HiddenImports["strconv"] = true
				converted = `__strconv.Itoa(` + expr + `)`
			case "f":
				template.HiddenImports["strconv"] = true
				converted = `__strconv.FormatFloat(` + expr + `, 'f', -1, 64)`
			case "s":
				converted = expr
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
	if template.Html.Len() > 0 {
		template.Buffer.WriteString(StartWrite + template.Html.String() + CloseWrite)
	}
	template.Html.Reset()
}

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

		if (token.Type == html.StartTagToken || token.Type == html.SelfClosingTagToken) && token.Data == "go" {
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

		} else if token.Type == html.EndTagToken && token.Data == "go" {
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
			text = template.replaceVariables(text)
			template.Html.WriteString(text)
		} else {
			// ----------------------------------------------------------------
			// Other tags, looking for $variable
			// ----------------------------------------------------------------

			tag := escapeBacktick(token.String())
			tag = template.replaceVariables(tag)
			template.Html.WriteString(tag)
		}
		previousToken = token
	}
	template.flush()
	return template
}

func (template *Template) ParseFile(file string) *Template {
	in, err := os.Open(file)
	if err != nil {
		log.Println("Unable to open", file)
		log.Fatal(err)
	}
	defer in.Close()
	return template.ParseReader(in)
}

func (template *Template) ParseString(data string) *Template {
	return template.ParseReader(strings.NewReader(data))
}
