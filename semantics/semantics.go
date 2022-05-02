package semantics

import (
	"Practica4/parser"
	"fmt"
)

// we have a postfix notation operation
// we want to create a DAG

type Semantics struct {
	Expression []parser.Token
}
type Node struct {
	op  string
	id1 int
	id2 int
}

func NewSemantics() *Semantics {
	return &Semantics{}
}

func (s *Semantics) Semantics() error {
	// iterate over the expression
	var aux interface{}
	var aux2 interface{}
	for i := 0; i < len(s.Expression); i++ {
		if s.Expression[i].Type == "number" {
			// we have to append a node to the register
			add(s.Expression[i])
			// if p1 has we have to modify p2
			if aux == nil {
				//fmt.Println("entered p1")
				aux = s.Expression[i]
			} else {
				// we have to modify p2
				aux2 = s.Expression[i]
			}
		} else if s.Expression[i].Type == "id" {
			add(s.Expression[i])
			if aux == nil {
				//fmt.Println("entered p1")
				aux = s.Expression[i]
			} else {
				// we have to modify p2
				aux2 = s.Expression[i]
			}
		} else if s.Expression[i].Type == "operator" {
			//fmt.Println("entered operator")
			//fmt.Println(aux)
			//fmt.Println(aux2)

			// get id1 and id2
			id1 := check(aux)
			id2 := check(aux2)
			//fmt.Println("ID1: ", id1)
			//fmt.Println("ID2: ", id2)
			if id2 == -1 {
				id2 = id1
				//id1 has to be the las NODE
				x, y := getLast()
				id1 = x

				// how do we make sense of the operation
				if id2 == -1 {
					id2 = y
					// if it's still -1 it is the same node
					if id2 == -1 {
						id2 = id1
					}
				}
			}
			// create a node
			node := Node{op: s.Expression[i].Info, id1: id1, id2: id2}
			// append the node to the register
			add(node)

			// reset aux and aux2
			aux = nil
			aux2 = nil
		}
	}
	fmt.Println(registers)
	return nil
}
