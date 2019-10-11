package types

import (
	"strings"
)

//CustomList has a list where each list seperated by comma ,
type CustomList struct {
	options []string
}

//NewCustomList is a constructor for CustomList
func NewCustomList(options string) CustomList {
	opt := strings.Split(options, ",")
	return CustomList{
		options: opt,
	}
}

//Options is a helper function that implements Options() from CustomList interface
func (c CustomList) Options() []string {
	return c.options
}
