package main
import (
	"bytes"
)
func Escape(buffer *bytes.Buffer) {
	buffer.WriteString(`<html>`)
	buffer.WriteString(`
    `)
	buffer.WriteString(`<body>`)
	buffer.WriteString(`
        `)
	buffer.WriteString(`<!-- in text -->`)
	buffer.WriteString(`
        test `+"`"+` test
        `)
	buffer.WriteString(`<!-- in attributes -->`)
	buffer.WriteString(`
        `)
	buffer.WriteString(`<div class="foo`+"`"+`bar" style="$$style">`)
	buffer.WriteString(`</div>`)
	buffer.WriteString(`
        `)
	buffer.WriteString(`<!-- in script -->`)
	buffer.WriteString(`
        `)
	buffer.WriteString(`<script>`)
	buffer.WriteString(`
            console.log("`+"`"+`");
        `)
	buffer.WriteString(`</script>`)
	buffer.WriteString(`
    `)
	buffer.WriteString(`</body>`)
	buffer.WriteString(`
`)
	buffer.WriteString(`</html>`)
}
