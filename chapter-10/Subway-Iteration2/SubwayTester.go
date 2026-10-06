package main

import (
	"fmt"
	"io"
)

func runSubwayTester(startName string, endName string, output io.Writer) error {
	loader := NewSubwayLoader()
	subway, err := loader.LoadFromFile("ObjectvilleSubway.txt")
	if err != nil {
		return err
	}

	if !subway.HasStation(startName) {
		return fmt.Errorf("%s は Objectville の駅ではありません", startName)
	}
	if !subway.HasStation(endName) {
		return fmt.Errorf("%s は Objectville の駅ではありません", endName)
	}

	route := subway.FindRoute(startName, endName)
	printer := NewSubwayPrinter(output)
	return printer.PrintDirections(route)
}
