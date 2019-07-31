package schematic

import (
	"go.lsl.digital/lardwaz/sdk/entity"
)

// Schematics is a collection of Schematic
type Schematics interface {
	// Filter returns a Schematics collection without elements filtered by fn (returning false)
	Filter(fn func(name string, item Schematic) bool) Schematics

	// Each loops over each item in the collection
	Each(fn func(name string, item Schematic))

	// Set schematic in collection
	Set(item Schematic)

	// Get a single item by name, if present
	Get(name string) (Schematic, bool)

	// Length returns the number of items in the collection
	Length() int
}

// Schematic represents metadata extracted from an entity Meta
type Schematic interface {
	// Name returns the name of the schematic item
	Name() string

	// Entity returns the entity represented by the schematic
	Entity() entity.Entity

	// Fields returns fields contained within the schematic
	Fields() Fields

	// Filters returns the filters for the schematic
	Filters() []string
}

// Fields represents a collection of fields attached to a schematic
type Fields interface {
	// Filter returns a Field collection without elements filtered by fn (returning false)
	Each(fn func(key string, item Field))

	// Each loops over each item in the collection
	Filter(fn func(key string, item Field) bool) Fields

	// Set sets a field item to the collecton
	Set(item Field)

	// Get a single field by name
	Get(name string) (Field, bool)

	// Length returns the number of items in the collection
	Length() int
}

// Field represents a single field belonging to a schematic
type Field interface {
	// Name returns the user-friendly machine name (typically based on json tag)
	Name() string

	// PropName returns the real property name (typically based on golang struct)
	PropName() string

	// Type is the datatype of the field (not necessarily golang based. see types.go)
	Type() string

	// Weight dictates the order of appearance, bigger weights "sink" and appear last
	Weight() int

	// Many indicates whether or not it contains multiple data entries
	Many() bool

	// Meta gives access to metadata associated with the field (typically json tags)
	Meta() FieldMeta
}

// FieldMeta represents a series of metadata associated with a field
type FieldMeta interface {
	// Allows to lookup a metadata value
	Lookup(key string) (value string, ok bool)
}
