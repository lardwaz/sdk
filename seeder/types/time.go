package types

//Time has the options of Time type
type Time struct {
	from int
	to   int
	val  string
}

//From is a helper function that implements Format() from Time interface
func (t Time) From() int {
	return t.from
}

//To is a helper function that implements Format() from Time interface
func (t Time) To() int {
	return t.to
}

//Val is a helper function that implements Format() from Time interface
func (t Time) Val() string {
	return t.val
}
