package types


type Seeded struct {
    fields map[string]interface{}
}

func (s *Seeded) SeedableFields() map[string]interface{} {
    return s.fields
}

func (s *Seeded) Add(name string, t interface{}) Seeded {
    s.fields[name] = t

    return *s
}