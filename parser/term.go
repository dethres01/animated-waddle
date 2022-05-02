package parser

import "fmt"

func (p *Parser) term() error {
	// T -> ( E ) | id | num
	// switch on lookahead token
	//fmt.Println("term", p.Lookahead_token, p.Lookahead_value)
	switch p.Lookahead_token {
	case "left parenthesis":
		// E -> ( E )
		// match left parenthesis
		err := p.Match("left parenthesis")
		if err != nil {
			return err
		}
		// check for E
		err = p.expr()
		if err != nil {
			return err
		}
		// match right parenthesis
		err = p.Match("right parenthesis")
		if err != nil {
			return err
		}
	case "id":
		// T -> id
		// match id
		// append to tokens
		p.Tokens = append(p.Tokens, p.Expression[p.Counter])
		fmt.Println("ID LEAF", p.Lookahead_token, p.Lookahead_value)
		// add to registers
		err := p.Match("id")
		if err != nil {
			return err
		}
	case "number":
		// T -> num
		// match num
		p.Tokens = append(p.Tokens, p.Expression[p.Counter])
		fmt.Println("NUMBER LEAF", p.Lookahead_token, p.Lookahead_value)
		err := p.Match("number")
		if err != nil {
			return err
		}
	}
	return nil
}
