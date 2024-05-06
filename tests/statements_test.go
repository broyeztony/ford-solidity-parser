package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatementList(t *testing.T) {

	program := `
	print("hello Ford!");
	42;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "arguments": [
          {
            "type": "StringLiteral",
            "value": "hello Ford!"
          }
        ],
        "callee": {
          "name": "print",
          "type": "Identifier"
        },
        "type": "CallExpression"
      },
      "type": "ExpressionStatement"
    },
    {
      "expression": {
        "type": "NumericLiteral",
        "value": 42
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
