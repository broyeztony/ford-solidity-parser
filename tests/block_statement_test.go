package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockStatement(t *testing.T) {

	program := `
	{
		42; 
		"Hello";
	}
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "body": [
        {
          "expression": {
            "type": "NumericLiteral",
            "value": 42
          },
          "type": "ExpressionStatement"
        },
        {
          "expression": {
            "type": "StringLiteral",
            "value": "Hello"
          },
          "type": "ExpressionStatement"
        }
      ],
      "type": "BlockStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
