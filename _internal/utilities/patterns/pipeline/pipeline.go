package pipeline

// Pipeline is a simple pipeline implementation.
type Pipeline[T any] struct {
	pipes []Pipe[T]
}

// NewPipeline creates a new pipeline with the given pipes.
func NewPipeline[T any](pipes []Pipe[T]) *Pipeline[T] {
	return &Pipeline[T]{pipes: pipes}
}

func (p *Pipeline[T]) Process(input T) T {
	for _, pipe := range p.pipes {
		input = pipe.Process(input)
	}
	return input
}
