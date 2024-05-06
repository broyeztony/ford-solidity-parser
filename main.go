package main

import (
	"fmt"
	parser "ford-solidity-parser/parser"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile("playground.ford")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	program := string(data)

	p := parser.NewParser(program)
	ast := p.Parse()

	fmt.Println(parser.Encode(ast))
}
