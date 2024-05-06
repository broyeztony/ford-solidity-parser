package tests

import (
	"ford-lang-parser/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyStatement(t *testing.T) {

	program := `
	;
	`

	p := parser.NewParser(program)
	ast := p.Parse()
	actual := parser.Encode(ast)

	expected := `{
  "body": [
    {
      "type": "EmptyStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
