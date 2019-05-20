package types

//CreditCard has related option need for CreditCard
type CreditCardNum struct {
	info string
}

//NewCreditCardNum is a constructor for CreditCardNum
func NewCreditCardNum(info string) CreditCardNum {
	return CreditCardNum{
		info: info,
	}
}

//Info is a helper function that implements Info() from CreditCard interface
func (i CreditCardNum) Info() string {
	return i.info
}
