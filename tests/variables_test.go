package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariableDeclaration(t *testing.T) {

	program := `
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
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}

func TestVariableAssignment(t *testing.T) {

	program := `
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
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
