// Code generated by Gotmpl
// DO NOT EDIT, I MEAN IT'S USELESS :)

package gotmpl

import (
	__bytes "bytes"
	"github.com/florentsolt/gotmpl/template/testdata/benchmark/model"
)

func SimpleQtc(u *model.User, buffer *__bytes.Buffer) {
	buffer.WriteString(`<html>
    <body>
        <h1>`)
	buffer.WriteString(u.FirstName)
	buffer.WriteString(`</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
            `)
	for _, colorName := range u.FavoriteColors {
		buffer.WriteString(`
                <li>`)
		buffer.WriteString(colorName)
		buffer.WriteString(`</li>
            `)
	}
	buffer.WriteString(`
        </ul>
    </body>
</html>`)
}
