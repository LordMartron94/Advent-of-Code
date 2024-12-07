package shared

type Direction struct {
	DeltaR int
	DeltaC int
}

func (d *Direction) TurnRight() Direction {
	return Direction{DeltaR: d.DeltaC, DeltaC: -d.DeltaR}
}
