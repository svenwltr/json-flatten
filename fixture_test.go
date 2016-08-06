package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"testing"

	"github.com/aryann/difflib"
)

var (
	updateGolden = flag.Bool("update", false, "update golden files")
)

func testReadFileString(t *testing.T, fname ...string) string {
	raw, err := ioutil.ReadFile(path.Join(fname...))
	if err != nil {
		t.Fatal(err.Error())
	}
	return string(raw)
}

func testReadFile(t *testing.T, fname ...string) io.Reader {
	r, err := os.Open(path.Join(fname...))
	if err != nil {
		t.Fatal(err.Error())
	}
	return r
}

func assertTextEquals(t *testing.T, expected, obtained string) {
	if expected != obtained {
		e := strings.Split(expected, "\n")
		o := strings.Split(obtained, "\n")

		output := bytes.Buffer{}
		output.WriteString("\n")

		idx := make([]int, 0)
		drs := difflib.Diff(e, o)

		for i, dr := range drs {
			if dr.Delta != difflib.Common {
				for j := i - 5; j <= i+5; j++ {
					idx = append(idx, j)
				}
			}
		}

		sort.Ints(idx)
		last := -1

		for _, i := range idx {
			if i < 0 {
				continue
			}

			if last == i {
				continue
			}

			if last+1 != i {
				output.WriteString("\n")
			}

			last = i

			output.WriteString(fmt.Sprintf("%5d: %s", i+1, drs[i].String()))
			output.WriteString("\n")
		}

		t.Error(output.String())
		t.Fatal("Output mismatches with golden file. " +
			"Use `go test -update` to update golden file.")
	}
}

func TestFixtures(t *testing.T) {
	const (
		DIR    = `./test-fixtures/`
		INPUT  = `-input.json`
		GOLDEN = `-golden.txt`
	)

	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), INPUT) {
			continue
		}

		basename := strings.TrimSuffix(file.Name(), INPUT)
		input := testReadFile(t, DIR, file.Name())

		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)

		app := &Application{
			Args: []string{"json-flatten"},
			Exit: func(exit int) {
			},
			Stdout: stdout,
			Stderr: stderr,
			Stdin:  input,
		}
		app.Run()

		if *updateGolden {
			ioutil.WriteFile(path.Join(DIR, basename+GOLDEN), stdout.Bytes(), 0644)
		} else {
			expected := testReadFileString(t, DIR, basename+GOLDEN)
			assertTextEquals(t, expected, stdout.String())
		}
	}

}
