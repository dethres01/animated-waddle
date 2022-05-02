package semantics

import "reflect"

// interface array for the D A G
var registers []interface{}

// func to check if content is already on array and return index
func check(content interface{}) int {
	for i, v := range registers {
		if v == content {
			return i
		}
	}
	return -1
}

// func to add content to array
func add(content interface{}) {
	// check if content is already on array
	// get index

	if check(content) == -1 {
		// not on array, add
		registers = append(registers, content)
	}
}

// grab the two last nodes from the register
func getLast() (int, int) {
	i1, i2 := -1, -1
	for i := len(registers) - 1; i >= 0; i-- {
		if reflect.TypeOf(registers[i]).String() == "semantics.Node" {
			if i1 == -1 {
				i1 = i
			} else {
				i2 = i
				break
			}
		}
	}
	return i1, i2
}
