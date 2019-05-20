package types

//ParagraphsN has related option need for ParagraphsN
type ParagraphsN struct {
	val int
}

//NewParagraphsN is a constructor for ParagraphsN
func NewParagraphsN(val int) ParagraphsN {
	return ParagraphsN{
		val: val,
	}
}

//Val is a helper function that implements Val() from ParagraphsN interface
func (p ParagraphsN) Val() int {
	return p.val
}
