// Code generated by Gotmpl
// DO NOT EDIT, I MEAN IT'S USELESS :)

package main
import (
	__bytes "bytes"
	__strconv "strconv"
	"strings"
	"fmt"
)
func Basic(title string, class string, flag bool, list []string, bar map[string]string, nbi int, nbf float64, buffer *__bytes.Buffer) {
	buffer.WriteString(`<html>

<head>
    <title></title>
</head>

<body>
    <h1>
        `)
	buffer.WriteString(strings.TrimSpace(title))
	buffer.WriteString(`
    </h1>
    <p class="`)
	buffer.WriteString(class)
	buffer.WriteString(`">
        `)
	if flag || false {
	buffer.WriteString(`
            <hr/> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
        `)
}
	buffer.WriteString(`
    </p>

    <ul>
        `)
	for index, item := range list {
	buffer.WriteString(`
            <li>
                `)
	buffer.WriteString(__strconv.Itoa(index))
	buffer.WriteString(` `)
	buffer.WriteString(item)
	buffer.WriteString(`
            </li>
        `)
}
	buffer.WriteString(`
    </ul>

    `)
	buffer.WriteString(`<pre><code>`)
	buffer.WriteString(fmt.Sprintf(`%#v`, bar))
	buffer.WriteString(`</code></pre>`)
	buffer.WriteString(`

    <script>
        var foo = 42;
        var nbi = `)
	buffer.WriteString(__strconv.Itoa(nbi))
	buffer.WriteString(`;
        var nbf = `)
	buffer.WriteString(__strconv.FormatFloat(nbf, 'f', -1, 64))
	buffer.WriteString(`;
        var hu = "`)
	buffer.WriteString(title)
	buffer.WriteString(`";
    </script>
</body>

</html>`)
}
