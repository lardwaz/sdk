package types

//DigitsN has the options of DigitsN
type DigitsN struct {
	val int
}

//NewDigitsN is a constructor for DigitsN
func NewDigitsN(val int) DigitsN {
	return DigitsN{
		val: val,
	}
}

//Val is a helper function that implements Val() from DigitsN s
func (d DigitsN) Val() int {
	return d.val
}
