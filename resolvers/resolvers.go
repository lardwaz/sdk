package entity

import (
	"github.com/graphql-go/graphql"
)

//Resolvable is an entity that can be resolved
type Resolvable interface {
	Resolvers() ResolverSet
}

//ResolverFactory denotes a function that generates a graphql resolver function
//this is used to generate listing functions given a dynamic filter list: on a map[field name] = filter type
type ResolverFactory func(schematic Schematic) graphql.FieldResolveFn

//ResolverSet provides a means to group all resolvers
type ResolverSet struct {
	Single  ResolverFactory
	Listing ResolverFactory
	Create  ResolverFactory
	Update  ResolverFactory
	Delete  ResolverFactory
}
