package main

import (
	"Practica4/parser"
	"Practica4/semantics"
	"fmt"
)

func main() {
	p := parser.NewParser()
	p.Expression = parser.ExpressionCatcher()
	err := p.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Expression)
	fmt.Println(p.Tokens)

	// semantic analysis
	s := semantics.NewSemantics()
	s.Expression = p.Tokens
	err = s.Semantics()
	if err != nil {
		fmt.Println(err)
	}

}
