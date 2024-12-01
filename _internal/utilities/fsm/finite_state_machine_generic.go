package fsm

import "context"

// Credits to: https://medium.com/@johnsiilver/go-state-machine-patterns-3b667f345b5e for the implementation.

type State[T any] func(ctx context.Context, args T) (T, State[T], error)

func Run[T any](ctx context.Context, args T, start State[T]) (T, error) {
	var err error
	current := start
	for {
		if ctx.Err() != nil {
			return args, ctx.Err()
		}
		args, current, err = current(ctx, args)
		if err != nil {
			return args, err
		}
		if current == nil {
			return args, nil
		}
	}
}
