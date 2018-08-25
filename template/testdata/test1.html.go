package testdata
import (
	"bytes"
	"strings"
	"strconv"
	"fmt"
)
func Test1(title string, class string, flag bool, list []string, bar map[string]string, buffer *bytes.Buffer) {
	buffer.WriteString(`
`)
	buffer.WriteString(`
`)
	buffer.WriteString(`

`)
	buffer.WriteString(`<html>`)
	buffer.WriteString(`

`)
	buffer.WriteString(`<head>`)
	buffer.WriteString(`
    `)
	buffer.WriteString(`<title>`)
	buffer.WriteString(`</title>`)
	buffer.WriteString(`
`)
	buffer.WriteString(`</head>`)
	buffer.WriteString(`

`)
	buffer.WriteString(`<body>`)
	buffer.WriteString(`
    `)
	buffer.WriteString(`<h1>`)
	buffer.WriteString(`
        `)
	buffer.WriteString(strings.TrimSpace(title))
	buffer.WriteString(`
    `)
	buffer.WriteString(`</h1>`)
	buffer.WriteString(`
    `)
	buffer.WriteString(`<p class="$class" $class="">`)
	buffer.WriteString(`
        `)
	if flag || false {
	buffer.WriteString(`
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
        `)
}
	buffer.WriteString(`
    `)
	buffer.WriteString(`</p>`)
	buffer.WriteString(`

    `)
	buffer.WriteString(`<ul>`)
	buffer.WriteString(`
        `)
	for index, item := range list {
	buffer.WriteString(`
            `)
	buffer.WriteString(`<li>`)
	buffer.WriteString(`
                `)
	buffer.WriteString(strconv.Itoa(index))
	buffer.WriteString(`
                `)
	buffer.WriteString(item)
	buffer.WriteString(`
            `)
	buffer.WriteString(`</li>`)
	buffer.WriteString(`
        `)
}
	buffer.WriteString(`
    `)
	buffer.WriteString(`</ul>`)
	buffer.WriteString(`

    Let&#39;s write a blackquote `)
	buffer.WriteString("`")
	buffer.WriteString(` for the fun :)

    `)
	buffer.WriteString(`<pre><code>`)
	buffer.WriteString(fmt.Sprintf(`%#v`, bar))
	buffer.WriteString(`</code></pre>`)
	buffer.WriteString(`

    `)
	buffer.WriteString(`<script>`)
	buffer.WriteString(`
        var foo = 42;
        var hu = foo;
    `)
	buffer.WriteString(`</script>`)
	buffer.WriteString(`
`)
	buffer.WriteString(`</body>`)
	buffer.WriteString(`

`)
	buffer.WriteString(`</html>`)
	buffer.WriteString(`

// includes // range`)
}
