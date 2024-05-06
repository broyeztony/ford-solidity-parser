package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariableDeclaration(t *testing.T) {

	program := `
	contract Playground;
	let x;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "declarations": [
        {
          "id": {
            "name": "x",
            "type": "Identifier"
          },
          "initializer": null,
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestVariableAssignment(t *testing.T) {

	program := `
	contract Playground;
	let x = 42;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "declarations": [
        {
          "id": {
            "name": "x",
            "type": "Identifier"
          },
          "initializer": {
            "type": "NumericLiteral",
            "value": 42
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
