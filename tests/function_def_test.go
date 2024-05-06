package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionDef(t *testing.T) {

	program := `
	contract Playground;

	def square {
    	return _.x * _.x;
	}
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
            "argument": {
              "left": {
                "computed": false,
                "object": {
                  "name": "_",
                  "type": "Identifier"
                },
                "property": {
                  "name": "x",
                  "type": "Identifier"
                },
                "type": "MemberExpression"
              },
              "operator": "*",
              "right": {
                "computed": false,
                "object": {
                  "name": "_",
                  "type": "Identifier"
                },
                "property": {
                  "name": "x",
                  "type": "Identifier"
                },
                "type": "MemberExpression"
              },
              "type": "BinaryExpression"
            },
            "type": "ReturnStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "name": {
        "name": "square",
        "type": "Identifier"
      },
      "type": "FunctionDeclaration"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
