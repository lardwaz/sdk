package seeder

import (
	"go.lsl.digital/lardwaz/sdk/seeder/types"
)

type Field struct {
	Name string
	Type interface{}
}

type Seedable struct {
	seedableFields map[string]interface{}
}

func (s Seedable) SeedableFields() map[string]interface{} {
	return s.seedableFields
}

func (s Seedable) New(fields ...Field) Seedable {

	for _, field := range fields {
		s.seedableFields[field.Name] = field.Type
	}

	return s
}

//Number is a helper function that return NumberI
func Number(min, max, val int) types.Number {
	return types.NewNumber(min, max, val)
}

//--------- CustomListI helper functions ----------//

//CustomList is a helper function that return CustomListI
func CustomList(list string) types.CustomList {
	return types.NewCustomList(list)
}

//--------- CreditCardI helper functions ----------//

//CreditCard is a helper function that return CreditCard
func CreditCard(info string) types.CreditCardNum {
	return types.NewCreditCardNum(info)
}

//Date is a helper function that return Date
func Date(from, to int, format string) types.Date {
	return types.NewDate(from, to, format)
}

//DateTime is a helper function that return DateTime
func DateTime(now bool, fromYY, toYY, fromTimeHH, toTimeHH int, format, timezone string) types.Datetime {
	return types.NewDatetime(now, fromYY, toYY, fromTimeHH, toTimeHH, format, timezone)
}

//Time is a helper function that return Time
func Time(from, to int, val string) types.Time {
	return types.NewTime(from, to, val)

}

//Password is a helper function that return Password
func Password(min, max int, allowUpper, allowNumeric, allowSpecialCharacters bool) types.Password {
	return types.NewPassword(min, max, allowUpper, allowNumeric, allowSpecialCharacters)
}

//Characters is a helper function that return CharactersN
func Characters(val int) types.CharactersN {
	return types.NewCharactersN(val)
}

//Digits is a helper function that return DigitsN
func Digits(val int) types.DigitsN {
	return types.NewDigitsN(val)
}

//Dob is a helper function that return Dob
func Dob(from, to int, format string) types.Dob {
	return types.NewDob(from, to, format)
}

//Paragraphs is a helper function that return ParagraphsN
func Paragraphs(val string) types.ParagraphsN {
	return types.NewParagraphsN(val)
}
