package shared

import "fmt"

type Direction struct {
	DeltaR int
	DeltaC int
}

func (d *Direction) TurnRight() Direction {
	return Direction{DeltaR: d.DeltaC, DeltaC: -d.DeltaR}
}

func (d *Direction) String() string {
	return fmt.Sprintf("(%d, %d)", d.DeltaR, d.DeltaC)
}
