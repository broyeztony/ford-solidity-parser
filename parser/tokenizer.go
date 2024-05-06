package parser

import (
	"fmt"
	"regexp"
)

type Rule struct {
	re        *regexp.Regexp
	tokenType string
}

type Tokenizer struct {
	input  string
	spec   []Rule
	cursor int
}

type Token struct {
	_type string
	value string
}

func NewTokenizer(input string) *Tokenizer {

	tokenizer := &Tokenizer{
		input:  input,
		cursor: 0,
	}

	// Tokenizer rules
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\s+`), tokenType: "WHITESPACE"})           // whitespace
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\/\/.*`), tokenType: "COMMENT"})           // single-line comment
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\/\*[\s\S]*?\*\/`), tokenType: "COMMENT"}) // multi-line comments
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^;`), tokenType: ";"})                      // delimiter
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^:`), tokenType: ":"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\{`), tokenType: "{"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\}`), tokenType: "}"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\(`), tokenType: "("})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\)`), tokenType: ")"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^,`), tokenType: ","})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\.`), tokenType: "."})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\[`), tokenType: "["})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\]`), tokenType: "]"})

	// ------------------------------------------------------------------------------------- KEYWORDS
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\blet\b`), tokenType: "let"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\bif\b`), tokenType: "if"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\belse\b`), tokenType: "else"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\btrue\b`), tokenType: "true"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\bfalse\b`), tokenType: "false"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\bnull\b`), tokenType: "null"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\bwhile\b`), tokenType: "while"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\bdo\b`), tokenType: "do"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\bfor\b`), tokenType: "for"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\bdef\b`), tokenType: "def"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\breturn\b`), tokenType: "return"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\brecover\b`), tokenType: "recover"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\bcontract\b`), tokenType: "contract"})

	// ------------------------------------------------------------------------------------- NUMBERS
	// TODO: https://golangbyexample.com/golang-regex-floating-point-number/
	// ^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)$
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)`), tokenType: "NUMBER"})

	// ------------------------------------------------------------------------------------- (currently) forbidden tokens
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\+\+`), tokenType: "FORBIDDEN_TOKEN"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\-\-`), tokenType: "FORBIDDEN_TOKEN"})

	// ------------------------------------------------------------------------------------- IDENTIFIER, OPERATORS
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\w+`), tokenType: "IDENTIFIER"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^[=!]=`), tokenType: "EQUALITY_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^->`), tokenType: "ERROR_HANDLER_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^=`), tokenType: "SIMPLE_ASSIGN"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^[\*\/\+\-]=`), tokenType: "COMPLEX_ASSIGN"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^[+\-]`), tokenType: "ADDITIVE_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^[*\/]`), tokenType: "MULTIPLICATIVE_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^[><]=?`), tokenType: "RELATIONAL_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^&&`), tokenType: "LOGICAL_AND"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^\|\|`), tokenType: "LOGICAL_OR"})
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^!`), tokenType: "LOGICAL_NOT"})

	// ------------------------------------------------------------------------------------- STRINGS
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^"[^"]*"`), tokenType: "STRING"}) // single quotes
	tokenizer.spec = append(tokenizer.spec, Rule{re: regexp.MustCompile(`^'[^']*'`), tokenType: "STRING"}) // double quotes

	// FindStringSubmatch

	return tokenizer
}

func (t *Tokenizer) hasMoreTokens() bool {
	return t.cursor < len(t.input)
}

func (t *Tokenizer) getNextToken() *Token {

	if !t.hasMoreTokens() {
		return nil
	}

	nextStr := t.input[t.cursor:len(t.input)]
	for _, rule := range t.spec {
		matches := rule.re.FindStringSubmatch(nextStr)

		if len(matches) == 0 {
			continue
		}

		t.cursor += len(matches[0])

		if rule.tokenType == "WHITESPACE" || rule.tokenType == "COMMENT" {
			return t.getNextToken()
		}

		if rule.tokenType == "FORBIDDEN_TOKEN" {
			panic(fmt.Sprintf("Invalid token: %v", matches[0]))
		}

		return &Token{
			_type: rule.tokenType,
			value: matches[0],
		}
	}

	panic(fmt.Sprintf("Unexpected token:%v", nextStr))
}

func (tok *Token) toString() string {
	return fmt.Sprintf("[TOKEN type: '%v', value: '%v']", tok._type, tok.value)
}
