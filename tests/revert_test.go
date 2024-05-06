package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRevert(t *testing.T) {

	program := `
	contract Playground;

	// evaluate division by zero will trigger the error handler block
	// revert will be executed and the program will terminate
	if balance < amount {
		revert();
	};
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)
	expected := `{
  "body": [
    {
      "alternate": null,
      "consequent": {
        "body": [
          {
            "expression": {
              "arguments": [],
              "callee": {
                "name": "revert",
                "type": "Identifier"
              },
              "type": "CallExpression"
            },
            "type": "ExpressionStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "test": {
        "left": {
          "name": "balance",
          "type": "Identifier"
        },
        "operator": "<",
        "right": {
          "name": "amount",
          "type": "Identifier"
        },
        "type": "BinaryExpression"
      },
      "type": "IfStatement"
    },
    {
      "type": "EmptyStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
