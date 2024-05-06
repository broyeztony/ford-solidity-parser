package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquality(t *testing.T) {

	program := `
	x > 0 == true;
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
        "operator": "==",
        "right": {
          "type": "BooleanLiteral",
          "value": true
        },
        "type": "BinaryExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
