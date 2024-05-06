package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoop(t *testing.T) {

	program := `
	do {
		x -= 1;
	} while (x > 10);
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "body": {
        "body": [
          {
            "expression": {
              "left": {
                "name": "x",
                "type": "Identifier"
              },
              "operator": "-=",
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
        "left": {
          "name": "x",
          "type": "Identifier"
        },
        "operator": ">",
        "right": {
          "type": "NumericLiteral",
          "value": 10
        },
        "type": "BinaryExpression"
      },
      "type": "DoWhileStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
