package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnaryExpression(t *testing.T) {

	program := `
	contract Playground;
	-x;
	!x;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "argument": {
          "name": "x",
          "type": "Identifier"
        },
        "operator": "-",
        "type": "UnaryExpression"
      },
      "type": "ExpressionStatement"
    },
    {
      "expression": {
        "argument": {
          "name": "x",
          "type": "Identifier"
        },
        "operator": "!",
        "type": "UnaryExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
