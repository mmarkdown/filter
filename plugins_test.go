package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mmarkdown/filter/renderer"
)

func TestPlugins(t *testing.T) {
	dir := "testdata"

	for name, impl := range Plugins {
		r := &renderer.R{}
		r.RegisterPlugin(impl)
		dir := path.Join(dir, name)

		testFiles, err := ioutil.ReadDir(dir)
		if err != nil {
			t.Fatalf("could not read %s: %q", dir, err)
		}

		for _, f := range testFiles {
			if f.IsDir() {
				continue
			}

			if filepath.Ext(f.Name()) != ".md" {
				continue
			}
			base := f.Name()[:len(f.Name())-3]

			t.Run(path.Join(dir, f.Name()), func(t *testing.T) {
				if err := doTestMarkdown(dir, base, r); err != nil {
					t.Errorf("failed test for %q: %s", base, err)
				}
			})
		}

	}

}

func doTestMarkdown(dir, basename string, r *renderer.R) error {
	filename := filepath.Join(dir, basename+".md")
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	filename = filepath.Join(dir, basename+".fmt")
	expected, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	expected = bytes.TrimSpace(expected)

	actual := r.Render(input)
	actual = bytes.TrimSpace(actual)

	if diff := cmp.Diff(string(actual), string(expected)); diff != "" {
		return fmt.Errorf("%s: differs: (-want +got)\n%s", basename+".md", diff)
	}

	return nil
}
