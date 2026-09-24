package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type SubwayLoader struct {
	subway *Subway
}

func NewSubwayLoader() *SubwayLoader {
	return &SubwayLoader{
		subway: NewSubway(),
	}
}

func (sl *SubwayLoader) LoadFromFile(filename string) (*Subway, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	err = sl.loadStations(sl.subway, reader)
	if err != nil {
		return nil, err
	}

	err = sl.loadLines(sl.subway, reader)
	if err != nil {
		return nil, err
	}

	return sl.subway, nil
}

func (sl *SubwayLoader) loadStations(subway *Subway, reader *bufio.Reader) error {
	for {
		currentLine, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		currentLine = strings.TrimSpace(currentLine)
		if len(currentLine) == 0 {
			break
		}
		subway.AddStation(currentLine)

		if err == io.EOF {
			break
		}
	}
	return nil
}

func (sl *SubwayLoader) loadLine(subway *Subway, reader *bufio.Reader, lineName string) error {
	station1Name, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}
	station1Name = strings.TrimSpace(station1Name)

	for {
		station2Name, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		station2Name = strings.TrimSpace(station2Name)
		if len(station2Name) == 0 {
			break
		}
		subway.AddConnection(station1Name, station2Name, lineName)
		station1Name = station2Name

		if err == io.EOF {
			break
		}
	}
	return nil
}

func (sl *SubwayLoader) loadLines(subway *Subway, reader *bufio.Reader) error {
	for {
		lineName, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		lineName = strings.TrimSpace(lineName)
		if len(lineName) == 0 {
			break
		}

		err = sl.loadLine(subway, reader, lineName)
		if err != nil {
			return err
		}

		if err == io.EOF {
			break
		}
	}
	return nil
}
