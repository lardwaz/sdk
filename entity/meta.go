package entity

//Meta represents meta data for an entity in the system
type Meta interface {
	Summary() string
	Instance() interface{}
}

//NewMeta creates a a new meta object
func NewMeta(summary string, instance interface{}) Meta {
	return meta{
		summary:  summary,
		instance: instance,
	}
}

type meta struct {
	summary  string
	instance interface{}
}

func (m meta) Summary() string {
	return m.summary
}

func (m meta) Instance() interface{} {
	return m.instance
}
