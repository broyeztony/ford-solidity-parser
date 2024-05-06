package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecover(t *testing.T) {

	program := `
	contract Playground;

	// evaluate division by zero will trigger the error handler block
	// recover will be executed and z will be assigned the value 1
	let z = x / 0 -> {
		recover 1;
	};
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)
	expected := `{
  "body": [
    {
      "declarations": [
        {
          "errorHandler": {
            "body": [
              {
                "argument": {
                  "type": "NumericLiteral",
                  "value": 1
                },
                "type": "RecoverStatement"
              }
            ],
            "type": "BlockStatement"
          },
          "id": {
            "name": "z",
            "type": "Identifier"
          },
          "initializer": {
            "left": {
              "name": "x",
              "type": "Identifier"
            },
            "operator": "/",
            "right": {
              "type": "NumericLiteral",
              "value": 0
            },
            "type": "BinaryExpression"
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}
