package template

import (
	"bytes"
	"fmt"
	gotemplate "html/template"
	"path"
	"path/filepath"
	"testing"

	"github.com/dchest/htmlmin"
	"github.com/florentsolt/gotmpl/template/testdata/benchmark/gotmpl"
	"github.com/florentsolt/gotmpl/template/testdata/benchmark/hero"
	"github.com/florentsolt/gotmpl/template/testdata/benchmark/model"
)

type tmplData struct {
	User     *model.User
	Nav      []*model.Navigation
	Title    string
	Messages []struct {
		I      int
		Plural bool
	}
}

var (
	testData = &model.User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
	}

	expectedtResult = `<html>
	<body>
		<h1>Bob</h1>
		<p>Here's a list of your favorite colors:</p>
		<ul>
			<li>blue</li>
			<li>green</li>
			<li>mauve</li>
		</ul>
	</body>
</html>`

	testComplexUser = &model.User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
		RawContent:     "<div><p>Raw Content to be displayed</p></div>",
		EscapedContent: "<div><div><div>Escaped</div></div></div>",
	}

	testComplexNav = []*model.Navigation{{
		Item: "Link 1",
		Link: "http://www.mytest.com/"}, {
		Item: "Link 2",
		Link: "http://www.mytest.com/"}, {
		Item: "Link 3",
		Link: "http://www.mytest.com/"},
	}
	testComplexTitle = testComplexUser.FirstName

	testComplexData = tmplData{

		User:  testComplexUser,
		Nav:   testComplexNav,
		Title: testComplexTitle,
		Messages: []struct {
			I      int
			Plural bool
		}{{1, false}, {2, true}, {3, true}, {4, true}, {5, true}},
	}

	expectedtComplexResult = `<!DOCTYPE html>
<html>
<body>
<header><title>Bob's Home Page</title>
<div class="header">Page Header</div>
</header>
<nav>
<ul class="navigation"><li><a href="http://www.mytest.com/">Link 1</a></li>
<li><a href="http://www.mytest.com/">Link 2</a></li>
<li><a href="http://www.mytest.com/">Link 3</a></li>
</ul>
</nav>
<section>
<div class="content">
	<div class="welcome">
			<h4>Hello Bob</h4>
			<div class="raw"><div><p>Raw Content to be displayed</p></div></div>
			<div class="enc">&lt;div&gt;&lt;div&gt;&lt;div&gt;Escaped&lt;/div&gt;&lt;/div&gt;&lt;/div&gt;</div>
	</div><p>Bob has 1 message</p><p>Bob has 2 messages</p><p>Bob has 3 messages</p><p>Bob has 4 messages</p><p>Bob has 5 messages</p>
</div>
</section>
<footer><div class="footer">copyright 2016</div>
</footer>
</body>
</html>`
)

/******************************************************************************
** Gotmpl
******************************************************************************/

func TestGotmpl(t *testing.T) {
	var buf bytes.Buffer

	gotmpl.SimpleQtc(testData, &buf)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkGotmpl(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		gotmpl.SimpleQtc(testData, &buf)
		buf.Reset()
	}
}

func TestComplexGotmpl(t *testing.T) {
	var buf bytes.Buffer

	gotmpl.Index(testComplexUser, testComplexNav, testComplexTitle, &buf)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexGotmpl(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		gotmpl.Index(testComplexUser, testComplexNav, testComplexTitle, &buf)
		buf.Reset()
	}
}

/******************************************************************************
** Go (html/template)
******************************************************************************/

func TestGolang(t *testing.T) {
	var buf bytes.Buffer

	tmpl, err := gotemplate.ParseFiles(path.Join("testdata", "benchmark", "go", "simple.tmpl"))
	if err != nil {
		t.Error(err)
	}
	err = tmpl.Execute(&buf, testData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkGolang(b *testing.B) {
	var buf bytes.Buffer

	tmpl, _ := gotemplate.ParseFiles(path.Join("testdata", "benchmark", "go", "simple.tmpl"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tmpl.Execute(&buf, testData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

func TestComplexGolang(t *testing.T) {

	var buf bytes.Buffer

	funcMap := gotemplate.FuncMap{
		"safehtml": func(text string) gotemplate.HTML { return gotemplate.HTML(text) },
	}

	templates := make(map[string]*gotemplate.Template)
	templatesDir := path.Join("testdata", "benchmark", "go")

	layouts, err := filepath.Glob(templatesDir + "/layout/*.tmpl")
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err)
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		templates[filepath.Base(layout)] = gotemplate.Must(gotemplate.New("").Funcs(funcMap).ParseFiles(files...))
	}
	templates["index.tmpl"].ExecuteTemplate(&buf, "base", testComplexData)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexGolang(b *testing.B) {
	var buf bytes.Buffer

	funcMap := gotemplate.FuncMap{
		"safehtml": func(text string) gotemplate.HTML { return gotemplate.HTML(text) },
	}

	templates := make(map[string]*gotemplate.Template)
	templatesDir := path.Join("testdata", "benchmark", "go")

	layouts, err := filepath.Glob(templatesDir + "/layout/*.tmpl")
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err)
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		templates[filepath.Base(layout)] = gotemplate.Must(gotemplate.New("").Funcs(funcMap).ParseFiles(files...))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		templates["index.tmpl"].ExecuteTemplate(&buf, "base", testComplexData)
		buf.Reset()
	}
}

/******************************************************************************
** Hero
******************************************************************************/
func TestHero(t *testing.T) {
	var buf bytes.Buffer

	hero.SimpleQtc(testData, &buf)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkHero(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		hero.SimpleQtc(testData, &buf)
		buf.Reset()
	}
}

func TestComplexHero(t *testing.T) {
	var buf bytes.Buffer

	hero.Index(testComplexUser, testComplexNav, testComplexTitle, &buf)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexHero(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		hero.Index(testComplexUser, testComplexNav, testComplexTitle, &buf)
		buf.Reset()
	}
}

/******************************************************************************
** helpers
******************************************************************************/

func linesEquals(str1, str2 string) (explanation string, equals bool) {
	if str1 == str2 {
		return "", true
	}

	// Minify removes whitespace infront of the first tag
	b1, err := htmlmin.Minify([]byte(str1), nil)
	if err != nil {
		panic(err)
	}

	b2, err := htmlmin.Minify([]byte(str2), nil)
	if err != nil {
		panic(err)
	}

	b1 = bytes.Replace(b1, []byte(" "), []byte("[space]"), -1)
	b1 = bytes.Replace(b1, []byte("\t"), []byte("[tab]"), -1)
	b1 = bytes.Replace(b1, []byte("\n"), []byte(""), -1)

	b2 = bytes.Replace(b2, []byte(" "), []byte("[space]"), -1)
	b2 = bytes.Replace(b2, []byte("\t"), []byte("[tab]"), -1)
	b2 = bytes.Replace(b2, []byte("\n"), []byte(""), -1)

	if bytes.Compare(b1, b2) != 0 {
		return fmt.Sprintf("Lines don't match \n1:\"%s\"\n2:\"%s\"", b1, b2), false
	}

	return "", true
}
