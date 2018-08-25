package template

import (
	"bytes"
	"strings"
)

func EscapeBackquote(value string, buffer *bytes.Buffer) {
	index := strings.IndexByte(value, '`')
	for index != -1 {
		buffer.WriteString(value[:index] + "`)\n")
		buffer.WriteString("\tbuffer.WriteString(\"`\")\n")
		buffer.WriteString("\tbuffer.WriteString(`")
		value = value[index+1 :]
		index = strings.IndexByte(value, '`')
	}
	buffer.WriteString(value)
}

// https://github.com/shiyanhui/hero/blob/2eee29f96f91f8ec1d177a69533507b62f1f9e83/util.go#L12
var (
	escapedKeys   = []byte{'&', '\'', '<', '>', '"'}
	escapedValues = []string{"&amp;", "&#39;", "&lt;", "&gt;", "&#34;"}
)

// EscapeHTML escapes the html and then put it to the buffer.
func EscapeHTML(html string, buffer *bytes.Buffer) {
	var i, j, k int

	for i < len(html) {
		for j = i; j < len(html); j++ {
			k = bytes.IndexByte(escapedKeys, html[j])
			if k != -1 {
				break
			}
		}

		buffer.WriteString(html[i:j])
		if k != -1 {
			buffer.WriteString(escapedValues[k])
		}
		i = j + 1
	}
}
