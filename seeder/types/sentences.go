package types

//SentencesN has related option need for SentencesN
type SentencesN struct {
	val int
}

//NewSentencesN is a constructor for SentencesN
func NewSentencesN(val int) SentencesN {
	return SentencesN{
		val: val,
	}
}

//Val is a helper function that implements Val() from SentencesN interface
func (s SentencesN) Val() int {
	return s.val
}
