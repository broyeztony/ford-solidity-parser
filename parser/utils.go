package parser

import (
	"bytes"
	"encoding/json"
	"strings"
)

func Encode(ast interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	encoder.Encode(ast)
	jsonString := buffer.String()
	return strings.TrimRight(jsonString, "\n")
}
