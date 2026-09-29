package main

import (
	"fmt"
	"io"
)

type SubwayPrinter struct {
	out io.Writer
}

func NewSubwayPrinter(out io.Writer) *SubwayPrinter {
	return &SubwayPrinter{out: out}
}

func (p *SubwayPrinter) PrintDirections(route []*Connection) error {
	if len(route) == 0 {
		_, err := fmt.Fprintln(p.out, "経路が見つかりませんでした。")
		return err
	}

	firstConnection := route[0]
	currentLine := firstConnection.GetLineName()
	if _, err := fmt.Fprintf(p.out, "出発駅は%sです。\n", firstConnection.GetStation1().GetName()); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(p.out, "%sに向かう%sに乗ります。\n", firstConnection.GetStation2().GetName(), currentLine); err != nil {
		return err
	}

	for _, connection := range route[1:] {
		if connection.GetLineName() == currentLine {
			if _, err := fmt.Fprintf(p.out, "	%sは通過します。\n", connection.GetStation1().GetName()); err != nil {
				return err
			}
			continue
		}

		if _, err := fmt.Fprintf(p.out, "%sに着いたら、%sを降ります。\n", connection.GetStation1().GetName(), currentLine); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(p.out, "%sに乗り換えて、%sに向かう電車に乗ります。\n", connection.GetLineName(), connection.GetStation2().GetName()); err != nil {
			return err
		}
		currentLine = connection.GetLineName()
	}

	lastConnection := route[len(route)-1]
	_, err := fmt.Fprintf(p.out, "%sに到着しました。ここで降りて、楽しんでください。\n", lastConnection.GetStation2().GetName())
	return err
}
