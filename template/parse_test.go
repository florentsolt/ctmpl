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
	// t.Fail()
}

func TestReplaceExpressions(t *testing.T) {
	template := New()

	cases := map[string]string{
		"lorem ipsum $s,foo$$s,foo$ lorem ipsum": "lorem ipsum `)\n\tbuffer.WriteString(foo)\n\tbuffer.WriteString(``)\n\tbuffer.WriteString(foo)\n\tbuffer.WriteString(` lorem ipsum",
		"lorem ipsum $$ lorem ipsum":             "lorem ipsum $ lorem ipsum",
		"lorem ipsum $s,foo$ lorem ipsum":        "lorem ipsum `)\n\tbuffer.WriteString(foo)\n\tbuffer.WriteString(` lorem ipsum",
	}
	for test, expected := range cases {
		result := template.replaceExpressions(test)
		if result != expected {
			t.Errorf("Result does not match, got %#v, expected %#v\n", result, expected)
		}
	}
}
