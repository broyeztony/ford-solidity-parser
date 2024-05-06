package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleBinaryExpression(t *testing.T) {

	program := `
	2 + 2;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 2
        },
        "operator": "+",
        "right": {
          "type": "NumericLiteral",
          "value": 2
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

func TestMulBinaryExpression(t *testing.T) {

	program := `
	2 + 2 * 2;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 2
        },
        "operator": "+",
        "right": {
          "left": {
            "type": "NumericLiteral",
            "value": 2
          },
          "operator": "*",
          "right": {
            "type": "NumericLiteral",
            "value": 2
          },
          "type": "BinaryExpression"
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

func TestParenthesizedBinaryExpression(t *testing.T) {

	program := `
	2 * (3 + 6);
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 2
        },
        "operator": "*",
        "right": {
          "left": {
            "type": "NumericLiteral",
            "value": 3
          },
          "operator": "+",
          "right": {
            "type": "NumericLiteral",
            "value": 6
          },
          "type": "BinaryExpression"
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
