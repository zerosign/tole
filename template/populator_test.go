package template

import (
	"github.com/aymerick/raymond/parser"
	"io/ioutil"
	"testing"
)

func TestPopulatorPopulateAst(t *testing.T) {

	bytes, err := ioutil.ReadFile("../test/database.yml.tmpl")

	if err != nil {
		t.Error(err)
	}

	ast, err := parser.Parse(string(bytes))

	if err != nil {
		t.Error(err)
	}

	populator := NewPopulator([]string{"alookup", "lookup", "rlookup", "hlookup"})

	_ = ast.Accept(populator)

	expected := SourcePaths{
		"test": []Operation{[]string{"alookup", "databases"}},
	}

	result := populator.Lists()

	if !result.Equal(expected) {
		t.Errorf("lookup paths expected to be %v but got %v", expected, result)
	}

}
