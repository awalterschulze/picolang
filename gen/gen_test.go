package gen

import (
	"bytes"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"testing"
)

func TestGen(t *testing.T) {
	funcs := map[string]string{
		"inc": "github.com/awalterschulze/picolang/funcs.Inc",
		"map": "github.com/awalterschulze/picolang/funcs.Map",
	}
	buf := bytes.NewBuffer(nil)
	err := Go(funcs, "127.0.0.1", buf)
	if err != nil {
		t.Error(err)
	}
	data := buf.Bytes()
	t.Logf("%v", (string(data)))
	os.Mkdir("./testdata/", 0777)
	if err := ioutil.WriteFile("./testdata/main.go", data, 0777); err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "./testdata/main.go", nil, parser.ImportsOnly); err != nil {
		t.Fatal(err)
	}
}
