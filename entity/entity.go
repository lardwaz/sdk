package entity

//Entity defines an object representation to be introspected
type Entity struct {
	instance interface{}
}

//Instance returns a instance of the entity at hand
func (e Entity) Instance() interface{} {
	return e.instance
}

//New creates a a new instance of entity given an object instance
func New(instance interface{}) Entity {
	return Entity{
		instance: instance,
	}
}
