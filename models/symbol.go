package models

import (
	"fmt"
	"strings"
)

type Symbol struct {
	fromSecurity Security
	toSecurity Security
}

func (s *Symbol) FromSecurity() Security {
	return s.fromSecurity
}

func (s *Symbol) ToSecurity() Security {
	return s.toSecurity
}

func (s *Symbol) String() string {
	return strings.ToLower(fmt.Sprintf("%s%s", s.fromSecurity.Name(), s.toSecurity.Name()))
}

func NewSymbol(fromSecurity Security, toSecurity Security) *Symbol {
	s := new(Symbol)
	s.fromSecurity = fromSecurity
	s.toSecurity = toSecurity
	return s
}
