package main

import (
	"fmt"
	"strings"
)

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

func (s *Subway) AddConnection(station1Name string, station2Name string, lineName string) {
	if s.HasStation(station1Name) && s.HasStation(station2Name) {
		station1 := NewStation(station1Name)
		station2 := NewStation(station2Name)
		connection := NewConnection(station1, station2, lineName)
		s.connections = append(s.connections, connection)
		reverseConnection := NewConnection(station2, station1, lineName)
		s.connections = append(s.connections, reverseConnection)
	} else {
		panic(fmt.Sprintf("Invalid connection: [%s, %s, %s]", station1Name, station2Name, lineName))
	}
}

func (s *Subway) HasConnection(station1Name string, station2Name string, lineName string) bool {
	station1 := NewStation(station1Name)
	station2 := NewStation(station2Name)
	for _, connection := range s.connections {
		if strings.EqualFold(connection.GetLineName(), lineName) {
			if connection.GetStation1().Equals(station1) && connection.GetStation2().Equals(station2) {
				return true
			}
		}
	}
	return false
}

/*
func main() {
	subway := NewSubway()

	fmt.Println(subway.HasStation("Ajax急流"))

	subway.AddStation("Ajax急流")

	fmt.Println(subway.HasStation("Ajax急流"))

	subway.AddStation("ajax急流")
	fmt.Println(len(subway.stations))

	loader := NewSubwayLoader()
	subway, err := loader.LoadFromFile("ObjectvilleSubway.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(len(subway.stations))
	fmt.Println(len(subway.connections))
}
*/
