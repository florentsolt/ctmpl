// Code generated by Gotmpl
// DO NOT EDIT, I MEAN IT'S USELESS :)

package gotmpl
import (
	__bytes "bytes"
)
func Header(title string, buffer *__bytes.Buffer) {
	buffer.WriteString(`<title>`)
	buffer.WriteString(title)
	buffer.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>
`)
}