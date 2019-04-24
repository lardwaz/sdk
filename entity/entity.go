package entity

//Entity defines an object representation to be introspected
type Entity interface {
	//Instance returns a instance of the entity at hand
	Instance() interface{}
}

type entity struct {
	instance interface{}
}

func (e entity) Instance() interface{} {
	return e.instance
}

//New creates a a new instance of entity given an object instance
func New(instance interface{}) Entity {
	return entity{
		instance: instance,
	}
}
