package main

import "fmt"

type Connection struct {
	station1 *Station
	station2 *Station
	lineName string
}

func NewConnection(station1 *Station, station2 *Station, lineName string) *Connection {
	return &Connection{
		station1: station1,
		station2: station2,
		lineName: lineName,
	}
}

func (c *Connection) GetStation1() *Station {
	return c.station1
}

func (c *Connection) GetStation2() *Station {
	return c.station2
}

func (c *Connection) GetLineName() string {
	return c.lineName
}

func (c *Connection) String() string {
	return fmt.Sprintf("[%s, %s, %s]", c.station1.GetName(), c.station2.GetName(), c.lineName)
}
