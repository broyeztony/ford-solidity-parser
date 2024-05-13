package main

import (
	"fmt"
	parser "ford-solidity-parser/parser"
	"io/ioutil"
)

func main() {

	filePath := "data/events.ford"
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	program := string(data)

	p := parser.NewParser(program)
	ast := p.Parse()

	fmt.Println(parser.Encode(ast))
}
