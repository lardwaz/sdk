package types

//WordsN has related option need for WordsN
type WordsN struct {
	val int
}

//NewWordsN is a constructor for WordsN
func NewWordsN(val int) WordsN {
	return WordsN{
		val: val,
	}
}

//Val is a helper function that implements Val() from WordsN interface
func (s WordsN) Val() int {
	return s.val
}
