package types

//DigitsN has the options of DigitsN
type DigitsN struct {
	val int
}

//Val is a helper function that implements Val() from DigitsN interface
func (d DigitsN) Val() int {
	return d.val
}
