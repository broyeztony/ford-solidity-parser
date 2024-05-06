package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorHandler(t *testing.T) {

	program := `
	let z = square({ x: a }) -> {};
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
              null
            ],
            "type": "BlockStatement"
          },
          "id": {
            "name": "z",
            "type": "Identifier"
          },
          "initializer": {
            "arguments": [
              {
                "type": "ObjectLiteral",
                "values": [
                  {
                    "name": "x",
                    "value": {
                      "name": "a",
                      "type": "Identifier"
                    }
                  }
                ]
              }
            ],
            "callee": {
              "name": "square",
              "type": "Identifier"
            },
            "type": "CallExpression"
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
