
package types

//Number has the options for Number type
type Number struct {
	min int
	max int
	val int
}


//Min is a helper function that implements Min() method from Number Interface
func (n Number) Min() int {
	return n.min
}

//NewNumber is a constructor for Number
func NewNumber(min, max, val int) Number {
	return  Number {
		min: min, max:max, val:val,
	}
}

//Max is a helper function that implements Max() method from Number Interface
func (n Number) Max() int {
	return n.max
}

//Val is a helper function that implements Val() method from Number Interface
func (n Number) Val() int {
	return n.val
}


// //SetMin is a helper function that implements SetMin() method from Number Interface
// func (n Number) SetMin(min int)  {
// 	n.min = min
// }

// //SetMax is a helper function that implements SetMax() method from Number Interface
// func (n Number) SetMax(max int)  {
// 	  n.max = max
// }

// //SetVal is a helper function that implements SetVal() method from Number Interface
// func (n Number) SetVal(val int) {
// 	n.val = val
// }