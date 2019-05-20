package types

//CharactersN has related option need for CharactersN
type CharactersN struct {
	val int
}

//NewCharactersN is a constructor for CharactersN
func NewCharactersN(val int) CharactersN {
	return CharactersN{
		val: val,
	}
}

//Val is a helper function that implements Val() from CharactersN interface
func (c CharactersN) Val() int {
	return c.val
}
