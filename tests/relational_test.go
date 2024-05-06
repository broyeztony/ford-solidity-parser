package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelationalExpression(t *testing.T) {

	program := `
	x > 0;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
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
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
