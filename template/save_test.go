package template

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
)

var cases = map[string]string{
	"basic": `
		package main
		import (
			"bytes"
			"os"
		)

		func main() {
			buffer := new(bytes.Buffer)
			Basic(
				"mytitle", 
				"myclass",
				true,
				[]string{"item1", "item2", "item3"},
				map[string]string{"k":"v"},
				42,
				42.42424242,
				buffer,
			)
			buffer.WriteTo(os.Stdout)
		}`,
	// ------------------------------------------------------------------------
	"escape": `
	package main
	import (
		"bytes"
		"os"
	)

	func main() {
		buffer := new(bytes.Buffer)
		Escape("<script></script>", buffer)
		buffer.WriteTo(os.Stdout)
	}`,
	// ------------------------------------------------------------------------
	"include": `
	package main
	import (
		"bytes"
		"os"
	)

	func main() {
		buffer := new(bytes.Buffer)
		Index(buffer)
		buffer.WriteTo(os.Stdout)
	}`,
}

func TestSave(t *testing.T) {
	for name, code := range cases {

		t.Run("Save "+name, func(t *testing.T) {
			files, err := ioutil.ReadDir(path.Join("testdata", name, "."))
			if err != nil {
				t.Error(err)
				return
			}
			for _, file := range files {
				if file.Name() == "out.html" || !strings.HasSuffix(file.Name(), ".html") {
					continue
				}
				template := New()
				template.ParseFile(path.Join("testdata", name, file.Name()))
				template.Save(path.Join("testdata", name, file.Name()+".go"))
			}

			out, err := os.Create(path.Join("testdata", name, "out.go"))
			if err != nil {
				t.Errorf("Unable to open out.go for %#v", name)
				t.Error(err)
			}
			defer out.Close()
			out.WriteString(code)
		})

		t.Run("Compile "+name, func(t *testing.T) {
			pkg := "github.com/florentsolt/gotmpl/template/testdata/" + name
			out := path.Join("testdata", name, name)
			cmd := exec.Command("go", "build", "-o", out, pkg)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("Failed to build %#v\n%s", name, output)
			}
		})

		t.Run("Execute "+name, func(t *testing.T) {
			cmd := exec.Command(path.Join("testdata", name, name))
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("Failed to run %#v\n%s", name, output)
				return
			}
			// Read out.html
		})

	}
	// t.Fail()
}

// Test Debug

// Test Trim
