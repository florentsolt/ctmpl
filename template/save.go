package template

import (
	"log"
	"os"
)

func (template *Template) Save(file string) {
	out, err := os.Create(file)
	if err != nil {
		log.Println("Unable to open", file)
		log.Fatal(err)
	}
	defer out.Close()

	out.WriteString("package " + template.Package + "\n")
	out.WriteString("import (\n")
	for name, _ := range template.Imports {
		out.WriteString("\t\"" + name + "\"\n")
	}
	out.WriteString(")\n")
	if template.FuncArgs == "" {
		out.WriteString("func " + template.FuncName + "(buffer *bytes.Buffer) {\n")
	} else {
		out.WriteString("func " + template.FuncName + "(" + template.FuncArgs + ", buffer *bytes.Buffer) {\n")
	}
	out.Write(template.Buffer.Bytes())
	out.WriteString("}\n")
}
