package main

import "fmt"

type Subway struct {
	stations    []*Station
	connections []*Connection
}

func NewSubway() *Subway {
	return &Subway{
		stations:    make([]*Station, 0),
		connections: make([]*Connection, 0),
	}
}

func (s *Subway) HasStation(stationName string) bool {
	target := NewStation(stationName)
	for _, station := range s.stations {
		if station.Equals(target) {
			return true
		}
	}
	return false
}

func (s *Subway) AddStation(stationName string) {
	if !s.HasStation(stationName) {
		station := NewStation(stationName)
		s.stations = append(s.stations, station)
	}
}

func main() {
	subway := NewSubway()

	fmt.Println(subway.HasStation("Ajax急流"))

	subway.AddStation("Ajax急流")

	fmt.Println(subway.HasStation("Ajax急流"))

	subway.AddStation("ajax急流")
	fmt.Println(len(subway.stations))
}
