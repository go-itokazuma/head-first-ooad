package main

import (
	"fmt"
	"strings"
)

type Subway struct {
	stations    map[int]*Station
	connections []*Connection
	network     map[int][]*Station
}

func NewSubway() *Subway {
	return &Subway{
		stations:    make(map[int]*Station),
		connections: make([]*Connection, 0),
		network:     make(map[int][]*Station),
	}
}

func (s *Subway) HasStation(stationName string) bool {
	return s.findStation(stationName) != nil
}

func (s *Subway) AddStation(stationName string) {
	station := NewStation(stationName)
	hash := station.HashCode()

	if _, exists := s.stations[hash]; !exists {
		s.stations[hash] = station
	}
}

func (s *Subway) findStation(stationName string) *Station {
	target := NewStation(stationName)
	hash := target.HashCode()

	station, exists := s.stations[hash]
	if !exists {
		return nil
	}

	if !station.Equals(target) {
		return nil
	}

	return station
}

func (s *Subway) AddConnection(station1Name string, station2Name string, lineName string) {
	if s.HasStation(station1Name) && s.HasStation(station2Name) {
		station1 := s.findStation(station1Name)
		station2 := s.findStation(station2Name)
		connection := NewConnection(station1, station2, lineName)
		s.connections = append(s.connections, connection)
		reverseConnection := NewConnection(station2, station1, lineName)
		s.connections = append(s.connections, reverseConnection)

		s.addToNetwork(station1, station2)

	} else {
		panic(fmt.Sprintf("Invalid connection: [%s, %s, %s]", station1Name, station2Name, lineName))
	}
}

func (s *Subway) HasConnection(station1Name string, station2Name string, lineName string) bool {
	station1 := s.findStation(station1Name)
	station2 := s.findStation(station2Name)
	for _, connection := range s.connections {
		if strings.EqualFold(connection.GetLineName(), lineName) {
			if connection.GetStation1().Equals(station1) && connection.GetStation2().Equals(station2) {
				return true
			}
		}
	}
	return false
}

func (s *Subway) addToNetwork(station1 *Station, station2 *Station) {
	hash1 := station1.HashCode()
	hash2 := station2.HashCode()

	s.network[hash1] = append(s.network[hash1], station2)
	s.network[hash2] = append(s.network[hash2], station1)
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
