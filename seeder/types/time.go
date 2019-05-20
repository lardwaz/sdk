package types

//Time has the options of Time type
type Time struct {
	from int
	to   int
	val  string
}

//NewTime is a constructor for Time
func NewTime(from, to int, val string) Time {
	return Time{
		from: from,
		to:   to,
		val:  val,
	}
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
