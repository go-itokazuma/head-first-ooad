package main

import "strings"

type Station struct {
	name string
}

func NewStation(name string) *Station {
	return &Station{
		name: name,
	}
}

func (s *Station) GetName() string {
	return s.name
}

// EqualsがtrueになるStationは、同じHashCodeを返さなければならない。
func (s *Station) Equals(other *Station) bool {
	if other == nil {
		return false
	}
	return strings.EqualFold(s.name, other.name)
}

func (s *Station) HashCode() int {
	hash := 0
	lowerName := strings.ToLower(s.name)
	for _, ch := range lowerName {
		hash = 31*hash + int(ch)
	}
	return hash
}
