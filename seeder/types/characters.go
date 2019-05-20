package types

//CharactersN has related option need for CharactersN
type CharactersN struct {
	val string
}

//NewCharactersN is a constructor for CharactersN
func NewCharactersN(val string) CharactersN {
	return CharactersN{
		val: val,
	}
}

//Val is a helper function that implements Val() from CharactersN interface
func (c CharactersN) Val() string {
	return c.val
}
