package types

// Seedable represents a seedable entity
type Seedable interface {
	SeedableFields() map[string]interface{}
	Multiplier() int
}

// Seeded implements Seedable
type Seeded struct {
	fields     map[string]interface{}
	multiplier int
}

// Field represents a seedable field
type Field struct {
	Name string
	Type interface{}
}

// NewSeedable returns a new Seedable
func NewSeedable(mul int, fields ...Field) Seedable {
	s := Seeded{multiplier: mul}

	s.fields = make(map[string]interface{})
	for _, field := range fields {
		s.fields[field.Name] = field.Type
	}

	return s
}

// SeedableFields returns a list of seedable fields
func (s Seeded) SeedableFields() map[string]interface{} {
	return s.fields
}

// Multiplier returns the multiplier for seeding
func (s Seeded) Multiplier() int {
	return s.multiplier
}
