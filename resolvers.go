package sdk

import "github.com/graphql-go/graphql"

//ResolverFactory denotes a functin that generates a graphql resolver function
//this is used to generate listing functions given a dynamic filter list: on a map[field name] = filter type
type ResolverFactory func(entity Meta) graphql.FieldResolveFn

//Resolvers provides a means to group all resolvers
type Resolvers struct {
	Single  ResolverFactory
	Listing ResolverFactory
	Create  ResolverFactory
	Update  ResolverFactory
	Delete  ResolverFactory
}
