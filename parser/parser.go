package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	tokenizer *Tokenizer
	lookahead *Token
}

func NewParser(input string) *Parser {

	parser := &Parser{}
	parser.tokenizer = NewTokenizer(strings.TrimSpace(input))

	return parser
}

func (p *Parser) Parse() interface{} {

	// to debug 👇🏻
	//for {
	//	token := p.tokenizer.getNextToken()
	//	fmt.Println("@ token", token)
	//
	//	if token == nil {
	//		break
	//	}
	//}

	p.lookahead = p.tokenizer.getNextToken()
	return p.Contract()
}

func (p *Parser) Program() interface{} {

	var body []interface{}

	body = p.StatementList("")

	program := map[string]interface{}{
		"type": "Program",
		"body": body,
	}

	return program
}

func (p *Parser) Contract() interface{} {

	p.eat("contract")
	id := p.Identifier()
	var contractName string
	if m, ok := id.(map[string]interface{}); ok {
		contractName = m["name"].(string)
	}
	p.eat(";")

	var contractBody []interface{}
	if p.lookahead != nil {
		contractBody = p.StatementList("")
	}

	return map[string]interface{}{
		"type": "Contract",
		"name": contractName,
		"body": contractBody,
	}
}

func (p *Parser) StatementList(stopLookahead string) []interface{} {
	statementList := []interface{}{p.Statement()}
	for p.lookahead != nil && p.lookahead._type != stopLookahead {
		statementList = append(statementList, p.Statement())
	}
	return statementList
}

func (p *Parser) Statement() interface{} {

	switch p.lookahead._type {
	case ";":
		return p.EmptyStatement()
	case "if":
		return p.IfStatement()
	case "{":
		return p.BlockStatement()
	case "let":
		return p.VariableStatement()
	case "def":
		return p.FunctionDeclaration()
	case "return":
		return p.ReturnStatement()
	case "recover":
		return p.RecoverStatement()
	case "while", "do", "for":
		return p.IterationStatement()
	default:
		return p.ExpressionStatement()
	}
}

func (p *Parser) EmptyStatement() interface{} {
	p.eat(";")
	return map[string]interface{}{
		"type": "EmptyStatement",
	}
}

func (p *Parser) ExpressionStatement() interface{} {

	expression := p.Expression()
	errorHandler := p.ErrorHandler()
	p.eat(";")

	node := map[string]interface{}{
		"type":       "ExpressionStatement",
		"expression": expression,
	}

	if errorHandler != nil {
		node["errorHandler"] = errorHandler
	}

	return node
}

func (p *Parser) ErrorHandler() interface{} {
	if p.lookahead._type == "ERROR_HANDLER_OPERATOR" {
		p.eat("ERROR_HANDLER_OPERATOR")
		handler := p.BlockStatement()

		return handler
	}
	return nil
}

func (p *Parser) IfStatement() interface{} {

	p.eat("if")

	test := p.Expression()
	consequent := p.BlockStatement()

	var alternate interface{}
	if p.lookahead != nil && p.lookahead._type == "else" {
		p.eat("else")
		alternate = p.BlockStatement()
	} else {
		alternate = nil
	}

	return map[string]interface{}{
		"type":       "IfStatement",
		"test":       test,
		"consequent": consequent,
		"alternate":  alternate,
	}
}

func (p *Parser) VariableStatement() interface{} {
	variableStatement := p.VariableStatementInit()
	p.eat(";")
	return variableStatement
}

func (p *Parser) VariableStatementInit() interface{} {
	p.eat("let")
	declarations := p.VariableDeclarationList()

	return map[string]interface{}{
		"type":         "VariableStatement",
		"declarations": declarations,
	}
}

func (p *Parser) VariableDeclarationList() interface{} {
	var declarations []interface{}
	for {
		declarations = append(declarations, p.VariableDeclaration())
		if !(p.lookahead._type == "," && p.eat(",") != nil) {
			break
		}
	}
	return declarations
}

func (p *Parser) VariableDeclaration() interface{} {

	id := p.Identifier()

	var initializer interface{}
	if p.lookahead._type != ";" && p.lookahead._type != "," {
		initializer = p.VariableInitializer()
	} else {
		initializer = nil
	}

	errorHandler := p.ErrorHandler()

	buffer := map[string]interface{}{
		"type":        "VariableDeclaration",
		"id":          id,
		"initializer": initializer,
	}

	if errorHandler != nil {
		buffer["errorHandler"] = errorHandler
	}

	return buffer
}

func (p *Parser) VariableInitializer() interface{} {
	p.eat("SIMPLE_ASSIGN")
	return p.AssignmentExpression()
}

func (p *Parser) Expression() interface{} {
	return p.AssignmentExpression()
}

// AssignmentExpression
// AssignmentExpression: LogicalORExpression | LHS '=' AssignmentExpression
func (p *Parser) AssignmentExpression() interface{} {

	left := p.LogicalORExpression()

	if !p.IsAssignmentOperator(p.lookahead._type) {
		return left
	}

	return map[string]interface{}{
		"type":     "AssignmentExpression",
		"operator": p.AssignmentOperator().value,
		"left":     p.CheckValidAssignmentTarget(left),
		"right":    p.AssignmentExpression(),
	}
}

func (p *Parser) IsAssignmentOperator(tokenType string) bool {
	return tokenType == "SIMPLE_ASSIGN" || tokenType == "COMPLEX_ASSIGN"
}

func (p *Parser) AssignmentOperator() *Token {
	if p.lookahead._type == "SIMPLE_ASSIGN" {
		return p.eat("SIMPLE_ASSIGN")
	}
	return p.eat("COMPLEX_ASSIGN")
}

func (p *Parser) CheckValidAssignmentTarget(node interface{}) interface{} {

	if inputMap, ok := node.(map[string]interface{}); ok {
		if inputMap["type"] == "Identifier" || inputMap["type"] == "MemberExpression" {
			return node
		}
	}
	panic("Invalid left-hand side in assignment expression")
}

func (p *Parser) LogicalORExpression() interface{} {
	return p.LogicalExpression("LogicalANDExpression", "LOGICAL_OR")
}

func (p *Parser) LogicalAndExpression() interface{} {
	return p.LogicalExpression("EqualityExpression", "LOGICAL_AND")
}

func (p *Parser) EqualityExpression() interface{} {
	return p.BinaryExpression("RelationalExpression", "EQUALITY_OPERATOR")
}

func (p *Parser) RelationalExpression() interface{} {
	return p.BinaryExpression("AdditiveExpression", "RELATIONAL_OPERATOR")
}

func (p *Parser) AdditiveExpression() interface{} {
	return p.BinaryExpression("MultiplicativeExpression", "ADDITIVE_OPERATOR")
}

func (p *Parser) MultiplicativeExpression() interface{} {
	return p.BinaryExpression("UnaryExpression", "MULTIPLICATIVE_OPERATOR")
}

func (p *Parser) LogicalExpression(builderName string, operatorToken string) interface{} {

	var left interface{}
	switch builderName {
	case "LogicalANDExpression":
		left = p.LogicalAndExpression()
	case "EqualityExpression":
		left = p.EqualityExpression()
	}

	for p.lookahead._type == operatorToken {
		operator := p.eat(operatorToken).value
		var right interface{}
		switch builderName {
		case "LogicalANDExpression":
			right = p.LogicalAndExpression()
		case "EqualityExpression":
			right = p.EqualityExpression()
		}

		left = map[string]interface{}{
			"type":     "LogicalExpression",
			"operator": operator,
			"left":     left,
			"right":    right,
		}
	}
	return left
}

func (p *Parser) BinaryExpression(builderName string, operatorToken string) interface{} {
	var left interface{}
	switch builderName {
	case "RelationalExpression":
		left = p.RelationalExpression()
	case "AdditiveExpression":
		left = p.AdditiveExpression()
	case "MultiplicativeExpression":
		left = p.MultiplicativeExpression()
	case "UnaryExpression":
		left = p.UnaryExpression()
	}

	for p.lookahead._type == operatorToken {
		operator := p.eat(operatorToken).value
		var right interface{}
		switch builderName {
		case "RelationalExpression":
			right = p.RelationalExpression()
		case "AdditiveExpression":
			right = p.AdditiveExpression()
		case "MultiplicativeExpression":
			right = p.MultiplicativeExpression()
		case "UnaryExpression":
			right = p.UnaryExpression()
		}

		left = map[string]interface{}{
			"type":     "BinaryExpression",
			"operator": operator,
			"left":     left,
			"right":    right,
		}
	}
	return left
}

func (p *Parser) UnaryExpression() interface{} {

	var operator string
	switch p.lookahead._type {
	case "ADDITIVE_OPERATOR":
		operator = p.eat("ADDITIVE_OPERATOR").value
		break
	case "LOGICAL_NOT":
		operator = p.eat("LOGICAL_NOT").value
		break
	}

	if operator != "" {
		return map[string]interface{}{
			"type":     "UnaryExpression",
			"operator": operator,
			"argument": p.UnaryExpression(),
		}
	}

	return p.LeftHandSideExpression()
}

func (p *Parser) LeftHandSideExpression() interface{} {
	return p.CallMemberExpression()
}

func (p *Parser) CallMemberExpression() interface{} {

	member := p.MemberExpression()
	if p.lookahead._type == "(" {
		return p.CallExpression(member)
	}
	return member
}

func (p *Parser) MemberExpression() interface{} {
	object := p.PrimaryExpression()

	for p.lookahead._type == "." || p.lookahead._type == "[" {

		if p.lookahead._type == "." {
			p.eat(".")
			property := p.Identifier()
			object = map[string]interface{}{
				"type":     "MemberExpression",
				"computed": false,
				"object":   object,
				"property": property,
			}
		}

		if p.lookahead._type == "[" {
			p.eat("[")
			property := p.Expression()
			p.eat("]")
			object = map[string]interface{}{
				"type":     "MemberExpression",
				"computed": true,
				"property": property,
				"object":   object,
			}
		}
	}

	return object
}

func (p *Parser) CallExpression(callee interface{}) interface{} {

	var callExpression interface{} = map[string]interface{}{
		"type":      "CallExpression",
		"callee":    callee,
		"arguments": p.Arguments(),
	}

	if p.lookahead._type == "(" {
		callExpression = p.CallExpression(callExpression)
	}

	return callExpression
}

func (p *Parser) Arguments() interface{} {
	p.eat("(")

	var argumentsList interface{}
	if p.lookahead._type != ")" {
		argumentsList = p.ArgumentList()
	} else {
		argumentsList = []interface{}{}
	}
	p.eat(")")
	return argumentsList
}

func (p *Parser) ArgumentList() interface{} {

	var argumentsList []interface{}
	for {
		argumentsList = append(argumentsList, p.AssignmentExpression())
		if !(p.lookahead._type == "," && p.eat(",") != nil) {
			break
		}
	}

	return argumentsList
}

func (p *Parser) PrimaryExpression() interface{} {

	if p.IsLiteral(p.lookahead._type) {
		return p.Literal()
	}

	switch p.lookahead._type {
	case "(":
		return p.ParenthesizedExpression()
	case "IDENTIFIER":
		return p.Identifier()
	default:
		return p.LeftHandSideExpression()
	}
}

func (p *Parser) ParenthesizedExpression() interface{} {
	p.eat("(")
	expression := p.Expression()
	p.eat(")")
	return expression
}

func (p *Parser) Identifier() interface{} {
	name := p.eat("IDENTIFIER").value
	return map[string]interface{}{
		"type": "Identifier",
		"name": name,
	}
}

func (p *Parser) IsLiteral(tokenType string) bool {
	return tokenType == "NUMBER" || tokenType == "STRING" || tokenType == "true" || tokenType == "false" || tokenType == "null" || tokenType == "{"
}

func (p *Parser) Literal() interface{} {

	switch p.lookahead._type {
	case "NUMBER":
		return p.NumericLiteral()
	case "STRING":
		return p.StringLiteral()
	case "true":
		return p.BooleanLiteral(true)
	case "false":
		return p.BooleanLiteral(false)
	case "null":
		return p.NullLiteral()
	case "{":
		return p.ObjectLiteral()
	}

	panic("Literal: unexpected literal production")
}

func (p *Parser) NumericLiteral() interface{} {
	token := p.eat("NUMBER")

	float, err := strconv.ParseFloat(token.value, 64)
	if err != nil {
		panic(fmt.Sprintf("Unexpected token value: '%v'", token.value))
	}

	return map[string]interface{}{
		"type":  "NumericLiteral",
		"value": float,
	}
}

func (p *Parser) StringLiteral() interface{} {
	token := p.eat("STRING")
	return map[string]interface{}{
		"type":  "StringLiteral",
		"value": token.value[1 : len(token.value)-1],
	}
}

func (p *Parser) BooleanLiteral(value bool) interface{} {

	if value {
		p.eat("true")
	} else {
		p.eat("false")
	}
	return map[string]interface{}{
		"type":  "BooleanLiteral",
		"value": value,
	}
}

func (p *Parser) NullLiteral() interface{} {
	p.eat("null")
	return map[string]interface{}{
		"type":  "NullLiteral",
		"value": "null",
	}
}

func (p *Parser) ObjectLiteral() interface{} {
	p.eat("{")
	var values []interface{}

	for p.lookahead != nil && p.lookahead._type != "}" {
		values = append(values, p.KeyValuePair())
	}
	p.eat("}")

	return map[string]interface{}{
		"type":   "ObjectLiteral",
		"values": values,
	}
}

func (p *Parser) KeyValuePair() interface{} {

	identifier := p.Identifier()
	var name string
	if m, ok := identifier.(map[string]interface{}); ok {
		name = m["name"].(string)
	}

	// handle shorthand notation. i.e. ```let x = { a };```
	var value interface{}
	if p.lookahead._type != ":" {
		value = map[string]interface{}{
			"type": "Identifier",
			"name": name,
		}

		if p.lookahead._type == "," {
			p.eat(",")
		}

		return map[string]interface{}{
			"name":  name,
			"value": value,
		}
	}

	p.eat(":")
	value = p.MemberExpression()
	if p.lookahead._type == "," {
		p.eat(",")
	}

	return map[string]interface{}{
		"name":  name,
		"value": value,
	}
}

func (p *Parser) BlockStatement() interface{} {

	p.eat("{")
	var body []interface{}
	if p.lookahead._type == "}" {
		// TODO: consider returning a null `body` in that case
		body = append(body, nil)
	} else {
		body = p.StatementList("}")
	}
	p.eat("}")

	return map[string]interface{}{
		"type": "BlockStatement",
		"body": body,
	}
}

func (p *Parser) FunctionDeclaration() interface{} {
	p.eat("def")
	name := p.Identifier()

	var params interface{}
	if p.lookahead._type == "(" {

		p.eat("(")
		if p.lookahead._type != ")" {
			params = p.FormalParameterList()
		} else {
			params = []interface{}{}
		}
		p.eat(")")
	}

	body := p.BlockStatement()
	buffer := map[string]interface{}{
		"type": "FunctionDeclaration",
		"name": name,
		"body": body,
	}

	if params != nil {
		buffer["params"] = params
	}

	return buffer
}

func (p *Parser) FormalParameterList() interface{} {
	var params []interface{}
	for {
		params = append(params, p.Identifier())
		if !(p.lookahead._type == "," && p.eat(",") != nil) {
			break
		}
	}

	return params
}

func (p *Parser) ReturnStatement() interface{} {
	p.eat("return")
	var argument interface{}
	if p.lookahead._type != ";" {
		argument = p.Expression()
	} else {
		argument = nil
	}
	p.eat(";")
	return map[string]interface{}{
		"type":     "ReturnStatement",
		"argument": argument,
	}
}

func (p *Parser) IterationStatement() interface{} {
	switch p.lookahead._type {
	case "while":
		return p.WhileStatement()
	case "do":
		return p.DoWhileStatement()
	case "for":
		return p.ForStatement()
	}
	return nil
}

func (p *Parser) WhileStatement() interface{} {
	p.eat("while")
	test := p.Expression()

	body := p.Statement()
	return map[string]interface{}{
		"type": "WhileStatement",
		"test": test,
		"body": body,
	}
}

func (p *Parser) DoWhileStatement() interface{} {
	p.eat("do")
	body := p.Statement()
	p.eat("while")
	test := p.Expression()
	p.eat(";")

	return map[string]interface{}{
		"type": "DoWhileStatement",
		"body": body,
		"test": test,
	}
}

func (p *Parser) ForStatement() interface{} {
	p.eat("for")
	p.eat("(")

	var init interface{}
	if p.lookahead._type == ";" {
		init = nil
	} else {
		init = p.ForStatementInit()
	}

	p.eat(";")

	var test interface{}
	if p.lookahead._type == ";" {
		test = nil
	} else {
		test = p.Expression()
	}

	p.eat(";")

	var update interface{}
	if p.lookahead._type == ")" {
		update = nil
	} else {
		update = p.Expression()
	}

	p.eat(")")

	body := p.Statement()

	return map[string]interface{}{
		"type":   "ForStatement",
		"init":   init,
		"test":   test,
		"update": update,
		"body":   body,
	}
}

func (p *Parser) ForStatementInit() interface{} {
	if p.lookahead._type == "let" {
		return p.VariableStatementInit()
	}
	return p.Expression()
}

func (p *Parser) RecoverStatement() interface{} {
	p.eat("recover")
	var argument interface{}
	if p.lookahead._type != ";" {
		argument = p.Expression()
	} else {
		argument = nil
	}

	p.eat(";")
	return map[string]interface{}{
		"type":     "RecoverStatement",
		"argument": argument,
	}
}

func (p *Parser) eat(tokenType string) *Token {

	token := p.lookahead

	if token == nil {
		panic(fmt.Sprintf("Unexpected end of input, expected: '%v'", tokenType))
	}
	if token._type != tokenType {
		panic(
			fmt.Sprintf("Unexpected token '%v', expected: '%v'", token.value, tokenType),
		)
	}

	p.lookahead = p.tokenizer.getNextToken()
	return token
}
