package template

import (
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func (template *Template) Parse(file string) {
	in, err := os.Open(file)
	if err != nil {
		log.Println("Unable to open", file)
		log.Fatal(err)
	}
	defer in.Close()

	tokenizer := html.NewTokenizer(in)

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
			log.Printf("Token %#v\n", token)
			if len(token.Attr) == 0 {
				continue
			}
			command, exists := template.Commands[token.Attr[0].Key]
			if !exists {
				log.Println("Command", token.Attr[0].Key, "does not exists, skipping...")
				continue
			}
			command(template, &token)
		} else if token.Type == html.EndTagToken && token.Data == "go" {
			template.Buffer.WriteString("}\n")
		} else {
			for _, attr := range token.Attr {
				if len(attr.Val) > 0 && attr.Val[0] == '$' {
					log.Printf("Attr %#v\n", attr)
				}
			}
			template.Buffer.WriteString("\tbuffer.WriteString(`")
			EscapeBackquote(token.String(), template.Buffer)
			template.Buffer.WriteString("`)\n")
			log.Println(token.String())
		}
	}
}