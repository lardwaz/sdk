package seeder

//NumberI has the options for Number type
type NumberI struct {
	min int
	max int
	val int
}

//CustomListI has a list where each list seperated by comma ,
type CustomListI struct {
	list string
}

//CreditCardI has related option need for CreditCard
type CreditCardI struct {
	info string
}

//DateI has the options of Date type
type DateI struct {
	from   int
	to     int
	format string
}

//DatetimeI has the options of Datetime type
type DatetimeI struct {
	now        bool
	fromYY     int
	toYY       int
	fromTimeHH int
	toTimeHH   int
	format     string
	timezone   string
}

//TimeI has the options of Time type
type TimeI struct {
	from int
	to   int
	val  string
}

//PasswordI has the options of Password type
type PasswordI struct {
	min                    int
	max                    int
	allowUpper             bool
	allowNumeric           bool
	allowSpecialCharacters bool
}

//--------- NumberI helper functions ----------//

//Min is a helper function that implements Min() method from Number Interface
func (n NumberI) Min() int {
	return n.min
}

//Max is a helper function that implements Max() method from Number Interface
func (n NumberI) Max() int {
	return n.max
}

//Val is a helper function that implements Val() method from Number Interface
func (n NumberI) Val() int {
	return n.val
}

//Number is a helper function that return NumberI
func Number(min, max, val int) NumberI {
	return NumberI{
		min: min, max: max, val: val,
	}
}

//--------- CustomListI helper functions ----------//

//List is a helper function that implements List() from CustomList interface
func (l CustomListI) List() string {
	return l.list
}

//CustomList is a helper function that return CustomListI
func CustomList(list string) CustomListI {
	return CustomListI{
		list: list,
	}
}

//--------- CreditCardI helper functions ----------//

//Info is a helper function that implements Info() from CreditCard interface
func (i CreditCardI) Info() string {
	return i.info
}

//CreditCard is a helper function that return CreditCardI
func CreditCard(info string) CreditCardI {
	return CreditCardI{
		info: info,
	}
}

//--------- DateI helper functions ----------//

//From is a helper function that implements From() from Date interface
func (d DateI) From() int {
	return d.from
}

//To is a helper function that implements To() from Date interface
func (d DateI) To() int {
	return d.to
}

//Format is a helper function that implements Format() from Date interface
func (d DateI) Format() string {
	return d.format
}

//Date is a helper function that return DateI
func Date(from, to int, format string) DateI {
	return DateI{
		from: from, to: to, format: format,
	}
}

//--------- DatetimeI helper functions ----------//

//Now is a helper function that implements Format() from DateTime interface
func (dt DatetimeI) Now() bool {
	return dt.now
}

//FromYY is a helper function that implements Format() from DateTime interface
func (dt DatetimeI) FromYY() int {
	return dt.fromYY
}

//ToYY is a helper function that implements Format() from DateTime interface
func (dt DatetimeI) ToYY() int {
	return dt.toYY
}

//FromTimeHH is a helper function that implements Format() from DateTime interface
func (dt DatetimeI) FromTimeHH() int {
	return dt.fromTimeHH
}

//ToTimeHH is a helper function that implements Format() from DateTime interface
func (dt DatetimeI) ToTimeHH() int {
	return dt.toTimeHH
}

//Format is a helper function that implements Format() from DateTime interface
func (dt DatetimeI) Format() string {
	return dt.format
}

//Timezone is a helper function that implements Format() from DateTime interface
func (dt DatetimeI) Timezone() string {
	return dt.timezone
}

//DateTime is a helper function that return DateTimeI
func DateTime(now bool, fromYY, toYY, fromTimeHH, toTimeHH int, format, timezone string) DatetimeI {
	return DatetimeI{
		now:        now,
		fromYY:     fromYY,
		toYY:       toYY,
		fromTimeHH: fromTimeHH,
		toTimeHH:   toTimeHH,
		format:     format,
		timezone:   timezone,
	}
}

//--------- Time helper functions ----------//

//From is a helper function that implements Format() from Time interface
func (t TimeI) From() int {
	return t.from
}

//To is a helper function that implements Format() from Time interface
func (t TimeI) To() int {
	return t.to
}

//Val is a helper function that implements Format() from Time interface
func (t TimeI) Val() string {
	return t.val
}

//Time is a helper function that return TimeI
func Time(from, to int, val string) TimeI {
	return TimeI{
		from: from,
		to:   to,
		val:  val,
	}
}

//--------- Password helper functions ----------//

// Min is a helper function that implements Format() from Password interface
func (p PasswordI) Min() int {
	return p.min
}

// Max is a helper function that implements Format() from Password interface
func (p PasswordI) Max() int {
	return p.max
}

// AllowUpper is a helper function that implements Format() from Password interface
func (p PasswordI) AllowUpper() bool {
	return p.allowUpper
}

// AllowNumeric is a helper function that implements Format() from Password interface
func (p PasswordI) AllowNumeric() bool {
	return p.allowNumeric
}

// AllowSpecialCharacters is a helper function that implements Format() from Password interface
func (p PasswordI) AllowSpecialCharacters() bool {
	return p.allowSpecialCharacters
}

//Password is a helper function that return PasswordI
func Password(min, max int, allowUpper, allowNumeric bool) PasswordI {
	return PasswordI{
		min:          min,
		max:          max,
		allowUpper:   allowUpper,
		allowNumeric: allowNumeric,
	}
}
