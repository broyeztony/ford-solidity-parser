package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleMemberExpression(t *testing.T) {

	program := `
	contract Playground;
	x.y;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "computed": false,
        "object": {
          "name": "x",
          "type": "Identifier"
        },
        "property": {
          "name": "y",
          "type": "Identifier"
        },
        "type": "MemberExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestMemberExpressionAssignment(t *testing.T) {

	program := `
	contract Playground;
	x.y = 1;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "computed": false,
          "object": {
            "name": "x",
            "type": "Identifier"
          },
          "property": {
            "name": "y",
            "type": "Identifier"
          },
          "type": "MemberExpression"
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
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestComputedExpression(t *testing.T) {

	program := `
	contract Playground;
	x[0] = 1;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "computed": true,
          "object": {
            "name": "x",
            "type": "Identifier"
          },
          "property": {
            "type": "NumericLiteral",
            "value": 0
          },
          "type": "MemberExpression"
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
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestComputedExpression2(t *testing.T) {

	program := `
	contract Playground;
	a.b.c['d'];
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "computed": true,
        "object": {
          "computed": false,
          "object": {
            "computed": false,
            "object": {
              "name": "a",
              "type": "Identifier"
            },
            "property": {
              "name": "b",
              "type": "Identifier"
            },
            "type": "MemberExpression"
          },
          "property": {
            "name": "c",
            "type": "Identifier"
          },
          "type": "MemberExpression"
        },
        "property": {
          "type": "StringLiteral",
          "value": "d"
        },
        "type": "MemberExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
