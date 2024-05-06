package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionCall(t *testing.T) {

	program := `
	contract Playground;

	foo(x)();
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "arguments": [],
        "callee": {
          "arguments": [
            {
              "name": "x",
              "type": "Identifier"
            }
          ],
          "callee": {
            "name": "foo",
            "type": "Identifier"
          },
          "type": "CallExpression"
        },
        "type": "CallExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestFunctionCall2(t *testing.T) {

	program := `
	contract Playground;

	console.log(x, y);
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
            "name": "x",
            "type": "Identifier"
          },
          {
            "name": "y",
            "type": "Identifier"
          }
        ],
        "callee": {
          "computed": false,
          "object": {
            "name": "console",
            "type": "Identifier"
          },
          "property": {
            "name": "log",
            "type": "Identifier"
          },
          "type": "MemberExpression"
        },
        "type": "CallExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
