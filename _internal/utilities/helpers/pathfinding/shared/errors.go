package shared

type OutOfBoundsError struct{}

func (e *OutOfBoundsError) Error() string {
	return "out of bounds"
}
