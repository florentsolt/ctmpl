package template

import (
	"testing"
)

var case1 = `
<html></html>
`

func TestParse(t *testing.T) {
	template := New()
	template.ParseString(case1)
	t.Fail()
}
