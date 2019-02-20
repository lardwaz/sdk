package sdk

//Meta represents a content type
type Meta interface {
	Description() string
	Instance() interface{}
}

//meta represents meta data for an entity in the system
type meta struct {
	description string
	instance    interface{}
	resolvers   Resolvers
}

//NewMeta create an instance of entity meta
func NewMeta(description string, instance interface{}, resolvers Resolvers) Meta {
	return meta{
		description: description,
		instance:    instance,
		resolvers:   resolvers,
	}
}

//Description method to implement meta interface
func (e meta) Description() string {
	return e.description
}

//Instance method to implement meta interface
func (e meta) Instance() interface{} {
	return e.instance
}

//Resolvers method to implement meta interface
func (e meta) Resolvers() Resolvers {
	return e.resolvers
}
