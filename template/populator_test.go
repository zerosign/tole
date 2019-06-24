package template

import (
	"github.com/aymerick/raymond/parser"
	"io/ioutil"
	"testing"
)

func TestPopulatorPopulateAst(t *testing.T) {

	bytes, err := ioutil.ReadFile("../test/database.yml.tmpl")

	t.Log("test")

	if err != nil {
		t.Error(err)
	}

	ast, err := parser.Parse(string(bytes))

	if err != nil {
		t.Error(err)
	}

	populator := NewPopulator()

	r := ast.Accept(populator)

	t.Logf("result: %v", r)

	t.Logf("variables: %v", populator.Lists())
}
