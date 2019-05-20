package types

//Year has the options of Year
type Year struct {
	min int
	max int
}

//NewYear is a constructor for Year
func NewYear(min, max int) Year {
	return Year{
		min: min, max: max,
	}
}

func (y Year) Min() int {
	return y.min
}

func (y Year) Max() int {
	return y.max
}
