package types

//Dob has the option for Dob
type Dob struct {
	from   int
	to     int
	format string
}

//NewDob is a constructor for Dob
func NewDob(from, to int, format string) Dob {
	return Dob{
		from:   from,
		to:     to,
		format: format,
	}
}

func (d Dob) From() int {
	return d.from
}

func (d Dob) To() int {
	return d.to
}

func (d Dob) Format() string {
	return d.format
}
