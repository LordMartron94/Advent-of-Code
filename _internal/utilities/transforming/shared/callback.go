package shared

import "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"

type TransformCallback[T comparable] func(node *shared.ParseTree[T])
