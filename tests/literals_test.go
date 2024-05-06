package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericLiteral(t *testing.T) {

	program := `
	contract Playground;
	42;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "type": "NumericLiteral",
        "value": 42
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestFloatLiteral(t *testing.T) {

	program := `
	contract Playground;
	42.34;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "type": "NumericLiteral",
        "value": 42.34
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestStringLiteral(t *testing.T) {

	program := `
	contract Playground;
	"42";
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "type": "StringLiteral",
        "value": "42"
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestSingleQuoteStringLiteral(t *testing.T) {

	program := `
	contract Playground;
	'"42"';
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "type": "StringLiteral",
        "value": "\"42\""
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestBooleanLiteral(t *testing.T) {

	program := `
	contract Playground;
	true;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "expression": {
        "type": "BooleanLiteral",
        "value": true
      },
      "type": "ExpressionStatement"
    }
  ],
  "name": "Playground",
  "type": "Contract"
}`

	assert.Equal(t, expected, actual)
}

func TestObjectLiteral(t *testing.T) {

	program := `
	contract Playground;

	let b = { x: a };
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "declarations": [
        {
          "id": {
            "name": "b",
            "type": "Identifier"
          },
          "initializer": {
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
