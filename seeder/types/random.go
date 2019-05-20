package types

//Random has the options of Random type
type Random struct {
	rtype string
}

//NewRandom is a constructor for Random
func NewRandom(rtype string) Random {
	return Random{
		rtype: rtype,
	}
}

// Rtype is a helper function that implements Format() from Random interface
func (r Random) Rtype() string {
	return r.rtype
}
