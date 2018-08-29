// Code generated by hero.
// source: /go/src/github.com/florentsolt/gotmpl/template/testdata/benchmark/hero/index.html
// DO NOT EDIT!
package hero

import (
	"bytes"

	"github.com/florentsolt/gotmpl/template/testdata/benchmark/model"
	"github.com/shiyanhui/hero"
)

func Index(u *model.User, nav []*model.Navigation, title string, buffer *bytes.Buffer) {
	buffer.WriteString(`

<!DOCTYPE html>
<html>
<body>

<header>
    `)
	buffer.WriteString(`<title>`)
	hero.EscapeHTML(title, buffer)
	buffer.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>
`)
	buffer.WriteString(`
</header>

<nav>
    `)
	buffer.WriteString(`<ul class="navigation">
    `)
	for _, item := range nav {
		buffer.WriteString(`
        <li><a href="`)
		hero.EscapeHTML(item.Link, buffer)
		buffer.WriteString(`">`)
		hero.EscapeHTML(item.Item, buffer)
		buffer.WriteString(`</a></li>
    `)
	}
	buffer.WriteString(`
</ul>`)
	buffer.WriteString(`
</nav>

<section>
<div class="content">
	<div class="welcome">
		<h4>Hello `)
	hero.EscapeHTML(u.FirstName, buffer)
	buffer.WriteString(`</h4>

		<div class="raw">`)
	buffer.WriteString(u.RawContent)
	buffer.WriteString(`</div>
		<div class="enc">`)
	hero.EscapeHTML(u.EscapedContent, buffer)
	buffer.WriteString(`</div>
	</div>

    `)
	for i := 1; i <= 5; i++ {
		if i == 1 {
			buffer.WriteString(`
            <p>`)
			hero.EscapeHTML(u.FirstName, buffer)
			buffer.WriteString(` has `)
			hero.FormatInt(int64(i), buffer)
			buffer.WriteString(` message</p>
        `)
		} else {
			buffer.WriteString(`
            <p>`)
			hero.EscapeHTML(u.FirstName, buffer)
			buffer.WriteString(` has `)
			hero.FormatInt(int64(i), buffer)
			buffer.WriteString(` messages</p>
        `)
		}
	}
	buffer.WriteString(`
</div>
</section>

<footer>
    `)
	buffer.WriteString(`<div class="footer">copyright 2016</div>`)
	buffer.WriteString(`
</footer>

</body>
</html>`)

}
