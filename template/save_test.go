package template

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
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
		Escape(buffer)
		buffer.WriteTo(os.Stdout)
	}`,
}

func TestSave(t *testing.T) {
	files, _ := ioutil.ReadDir("testdata")
	for _, file := range files {
		if !file.IsDir() || file.Name()[0] == '.' {
			continue
		}
		t.Run("Save "+file.Name(), func(t *testing.T) {
			template := New()
			template.ParseFile(path.Join("testdata", file.Name(), "in.html"))
			template.Save(path.Join("testdata", file.Name(), "in.html.go"))

			out, err := os.Create(path.Join("testdata", file.Name(), "out.go"))
			if err != nil {
				t.Errorf("Unable to open out.go for %#v", file.Name())
				t.Error(err)
			}
			defer out.Close()
			out.WriteString(cases[file.Name()])
		})

		t.Run("Compile "+file.Name(), func(t *testing.T) {
			pkg := "github.com/florentsolt/gotmpl/template/testdata/" + file.Name()
			out := path.Join("testdata", file.Name(), file.Name())
			cmd := exec.Command("go", "build", "-o", out, pkg)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("Failed to build %#v\n%s", file.Name(), output)
			}
		})

		t.Run("Execute "+file.Name(), func(t *testing.T) {
			cmd := exec.Command(path.Join("testdata", file.Name(), file.Name()))
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("Failed to run %#v\n%s", file.Name(), output)
				return
			}
			// Read out.html
		})

	}
	t.Fail()
}

// Test Debug

// Test Trim
