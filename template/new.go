package template

import (
	"bytes"
)

func New() *Template {
	return &Template{
		Imports:  map[string]bool{"bytes": true},
		Buffer:   new(bytes.Buffer),
		Commands: DefaultCommands,
		SilentTags: map[string]bool{
			"package": true,
			"import":  true,
			"func":    true,
		},
		Trim:  false,
		Debug: false,
	}
}
