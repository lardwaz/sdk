package entity

//Meta represents meta data for an entity in the system
type Meta interface {
	Info() string
	Instance() interface{}
	Resolvers() Resolvers
}

//NewMeta creates a a new meta object
func NewMeta(info string, instance interface{}, resolvers Resolvers) Meta {
	return meta{
		info:      info,
		instance:  instance,
		resolvers: resolvers,
	}
}

type meta struct {
	info      string
	instance  interface{}
	resolvers Resolvers
}

func (m meta) Info() string {
	return m.info
}

func (m meta) Instance() interface{} {
	return m.instance
}

func (m meta) Resolvers() Resolvers {
	return m.resolvers
}
