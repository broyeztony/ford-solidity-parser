package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfElse(t *testing.T) {

	program := `
	if x {
		x = 1;
	}
	else {
		x = 2;
	}
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "alternate": {
        "body": [
          {
            "expression": {
              "left": {
                "name": "x",
                "type": "Identifier"
              },
              "operator": "=",
              "right": {
                "type": "NumericLiteral",
                "value": 2
              },
              "type": "AssignmentExpression"
            },
            "type": "ExpressionStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "consequent": {
        "body": [
          {
            "expression": {
              "left": {
                "name": "x",
                "type": "Identifier"
              },
              "operator": "=",
              "right": {
                "type": "NumericLiteral",
                "value": 1
              },
              "type": "AssignmentExpression"
            },
            "type": "ExpressionStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "test": {
        "name": "x",
        "type": "Identifier"
      },
      "type": "IfStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
