package parser

func (p *Parser) expr() error {
	// E -> T E'
	// check for T
	//fmt.Println("expr", p.Lookahead_token, p.Lookahead_value)
	err := p.term()
	if err != nil {
		return err
	}
	// check for E'
	err = p.rest_of_expr()
	if err != nil {
		return err
	}
	return nil
}
