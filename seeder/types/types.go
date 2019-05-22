package types

type Seeded struct {
	fields map[string]interface{}
}

type Field struct {
	Name string
	Type interface{}
}

func NewSeeded(fields ...Field) Seeded {

	s := Seeded{}

	s.fields = make(map[string]interface{})
	for _, field := range fields {

		s.fields[field.Name] = field.Type
	}

	return s
}

func (s Seeded) SeedableFields() map[string]interface{} {
	return s.fields
}

func (s Seeded) Add(name string, t interface{}) Seeded {
	s.fields[name] = t

	return s
}
