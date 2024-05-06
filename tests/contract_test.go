package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContract(t *testing.T) {

	program := `
	contract Playground;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": null,
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
