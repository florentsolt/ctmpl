package template

import (
	"bytes"
)

// New returns a fresh template ready use with default values
func New() *Template {
	return &Template{
		Imports:       map[string]bool{},
		HiddenImports: map[string]bool{"bytes": true},
		Buffer:        new(bytes.Buffer),
		Commands:      DefaultCommands,
		SilentTags: map[string]bool{
			"package": true,
			"import":  true,
			"func":    true,
		},
		Trim:  false,
		Debug: false,
		Expr:  "$",
		Tag:   "go",
	}
}
