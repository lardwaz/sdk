package entity

//Meta represents meta data for an entity in the system
type Meta interface {
	Summary() string
	Instance() interface{}
	Resolvers() Resolvers
}

//NewMeta creates a a new meta object
func NewMeta(summary string, instance interface{}, resolvers Resolvers) Meta {
	return meta{
		summary:   summary,
		instance:  instance,
		resolvers: resolvers,
	}
}

type meta struct {
	summary   string
	instance  interface{}
	resolvers Resolvers
}

func (m meta) Summary() string {
	return m.summary
}

func (m meta) Instance() interface{} {
	return m.instance
}

func (m meta) Resolvers() Resolvers {
	return m.resolvers
}
