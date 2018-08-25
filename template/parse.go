package template

import (
	"io"
	"log"
	"os"
	"strings"
	"regexp"

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

func replaceVariables(value string) string {
	re := regexp.MustCompile(`([^\$])\$([\w\d\-\_]+)`)
	return re.ReplaceAllString(value, "$1`)\n\tbuffer.WriteString($2)\n\tbuffer.WriteString(`")
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
			command(template, &token)

		} else if token.Type == html.EndTagToken && token.Data == "go" {
			// ----------------------------------------------------------------
			// Close </go>
			// ----------------------------------------------------------------

			template.debug(&token)
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
			template.Buffer.WriteString("\tbuffer.WriteString(`")
			text := escapeBacktick(token.Data)
			if template.Trim {
				text = strings.TrimSpace(text)
			}
			template.Buffer.WriteString(text)
			template.Buffer.WriteString("`)\n")
		} else {
			// ----------------------------------------------------------------
			// Other tags, looking for $variable
			// ----------------------------------------------------------------

			tag := escapeBacktick(token.String())
			tag = replaceVariables(tag)
			template.Buffer.WriteString("\tbuffer.WriteString(`")
			template.Buffer.WriteString(tag)
			template.Buffer.WriteString("`)\n")
		}
		previousToken = token
	}
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
