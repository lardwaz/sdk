package entity

//Meta represents meta data for an entity in the system
type Meta struct {
	Description string
	Instance    interface{}
	Resolvers   Resolvers
}
