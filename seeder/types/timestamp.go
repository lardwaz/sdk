package types

//Timestamp has the options of Timestamp
type Timestamp struct {
	now        bool
	fromYY     int
	toYY       int
	fromTimeHH int
	toTimeHH   int
	format     string
	timezone   string
}

//NewTimestamp is a constructor for Timestamp
func NewTimestamp(now bool, fromYY int, toYY int, fromTimeHH int, toTimeHH int, format string, timezone string) Timestamp {

	return Timestamp{
		now:        now,
		fromYY:     fromYY,
		toYY:       toYY,
		fromTimeHH: fromTimeHH,
		toTimeHH:   toTimeHH,
		format:     format,
		timezone:   timezone,
	}
}

func (t Timestamp) Now() bool {
	return t.now
}

func (t Timestamp) FromYY() int {
	return t.fromYY
}

func (t Timestamp) ToYY() int {
	return t.toYY
}

func (t Timestamp) FromTimeHH() int {
	return t.fromTimeHH
}

func (t Timestamp) ToTimeHH() int {
	return t.toTimeHH
}

func (t Timestamp) Format() string {
	return t.format
}

func (t Timestamp) Timezone() string {
	return t.timezone
}
