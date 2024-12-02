package pipeline

// Pipe is an interface for pipeline operations.
type Pipe[T any] interface {
	Process(input T) T
}
