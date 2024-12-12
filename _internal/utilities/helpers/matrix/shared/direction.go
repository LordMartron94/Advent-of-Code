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

type DirectionExternal int

const (
	Down DirectionExternal = iota
	Left
	Right
	Up
)

func (d DirectionExternal) ToDirection() Direction {
	switch d {
	case Down:
		return Direction{DeltaR: 1, DeltaC: 0}
	case Left:
		return Direction{DeltaR: 0, DeltaC: -1}
	case Right:
		return Direction{DeltaR: 0, DeltaC: 1}
	case Up:
		return Direction{DeltaR: -1, DeltaC: 0}
	default:
		return Direction{}
	}
}
