package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogicalExpression(t *testing.T) {

	program := `
	contract Playground;
	x > 0 && y < 1;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "left": {
            "name": "x",
            "type": "Identifier"
          },
          "operator": ">",
          "right": {
            "type": "NumericLiteral",
            "value": 0
          },
          "type": "BinaryExpression"
        },
        "operator": "&&",
        "right": {
          "left": {
            "name": "y",
            "type": "Identifier"
          },
          "operator": "<",
          "right": {
            "type": "NumericLiteral",
            "value": 1
          },
          "type": "BinaryExpression"
        },
        "type": "LogicalExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
