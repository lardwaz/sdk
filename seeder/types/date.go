package types

//Date has the options of Date type
type Date struct {
	from   int
	to     int
	format string
}

//NewDate is a constructor for Date
func NewDate(from, to int, format string) Date {
	return Date{
		from:   from,
		to:     to,
		format: format,
	}
}

//From is a helper function that implements From() from Date interface
func (d Date) From() int {
	return d.from
}

//To is a helper function that implements To() from Date interface
func (d Date) To() int {
	return d.to
}

//Format is a helper function that implements Format() from Date interface
func (d Date) Format() string {
	return d.format
}
