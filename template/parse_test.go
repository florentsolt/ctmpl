package template

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	template := New()
	template.Parse("testdata/test1.html")
	template.Save("testdata/test1.html.go")
	t.Fail()
}
