package types

//Datetime has the options of Datetime type
type Datetime struct {
	now        bool
	fromYY     int
	toYY       int
	fromTimeHH int
	toTimeHH   int
	format     string
	timezone   string
}

//NewDatetime is a constructor for Datetime
func NewDatetime(now bool, fromYY int, toYY int, fromTimeHH int, toTimeHH int, format string, timezone string) Datetime {

	return Datetime{
		now:        now,
		fromYY:     fromYY,
		toYY:       toYY,
		fromTimeHH: fromTimeHH,
		toTimeHH:   toTimeHH,
		format:     format,
		timezone:   timezone,
	}
}

//Now is a helper function that implements Format() from DateTime interface
func (dt Datetime) Now() bool {
	return dt.now
}

//FromYY is a helper function that implements Format() from DateTime interface
func (dt Datetime) FromYY() int {
	return dt.fromYY
}

//ToYY is a helper function that implements Format() from DateTime interface
func (dt Datetime) ToYY() int {
	return dt.toYY
}

//FromTimeHH is a helper function that implements Format() from DateTime interface
func (dt Datetime) FromTimeHH() int {
	return dt.fromTimeHH
}

//ToTimeHH is a helper function that implements Format() from DateTime interface
func (dt Datetime) ToTimeHH() int {
	return dt.toTimeHH
}

//Format is a helper function that implements Format() from DateTime interface
func (dt Datetime) Format() string {
	return dt.format
}

//Timezone is a helper function that implements Format() from DateTime interface
func (dt Datetime) Timezone() string {
	return dt.timezone
}
