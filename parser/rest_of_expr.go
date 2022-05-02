package parser

import "fmt"

func (p *Parser) rest_of_expr() error {

	// switch case of token
	// E' -> + T E' | - T E' | e
	//fmt.Println("rest_of_expr", p.Lookahead_token, p.Lookahead_value)
	switch p.Lookahead_token {
	case "operator":
		// match operator
		aux := p.Expression[p.Counter]
		err := p.Match("operator")
		if err != nil {
			return err
		}
		// we can either have a term or an id
		err = p.term()
		if err != nil {
			return err
		}
		p.Tokens = append(p.Tokens, aux)
		fmt.Println("OP NODE", aux)
		err = p.rest_of_expr()
		if err != nil {
			return err
		}
	case "right parenthesis":
	case "left parenthesis":
	case "epsilon":
		// done
	}
	return nil
}
