package models

type Security struct {
	name string
}

func (s *Security) Name() string {
	return s.name
}

func NewSecurity(name string) *Security {
	s := new(Security)
	s.name = name
	return s
}
