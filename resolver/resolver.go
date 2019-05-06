package resolver

const (
	//Single resolvable for entity
	Single uint8 = 1

	//Listing resolvable for entity
	Listing uint8 = 2

	//Create resolvable for entity
	Create uint8 = 4

	//Update resolvable for entity
	Update uint8 = 8

	//Delete resolvable for entity
	Delete uint8 = 16

	//ALL resolve to all
	ALL uint8 = Single | Listing | Create | Update | Delete
)

//Resolvable indicates if an entity can be resolved using graphql
//and which methods to allow
type Resolvable interface {
	ResolverSummary() string
	SingleResolver() bool
	ListingResolver() bool
	CreateResolver() bool
	UpdateResolver() bool
	DeleteResolver() bool
}

type resolvable struct {
	summary string
	single  bool
	listing bool
	create  bool
	update  bool
	delete  bool
}

func (r resolvable) ResolverSummary() string {
	return r.summary
}

func (r resolvable) SingleResolver() bool {
	return r.single
}

func (r resolvable) ListingResolver() bool {
	return r.listing
}

func (r resolvable) CreateResolver() bool {
	return r.create
}

func (r resolvable) UpdateResolver() bool {
	return r.update
}

func (r resolvable) DeleteResolver() bool {
	return r.delete
}

//New resolvable entity
func New(summary string, resolvers uint8) Resolvable {
	return resolvable{
		summary: summary,
		single:  resolvers&Single != 0,
		listing: resolvers&Listing != 0,
		create:  resolvers&Create != 0,
		update:  resolvers&Update != 0,
		delete:  resolvers&Delete != 0,
	}
}
