package graph

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
)

// GraphHelper is a helper struct for working with a graph data structure.
// Inspired by https://www.geeksforgeeks.org/topological-sorting/
type GraphHelper[T any] struct {
	matrixHelper *matrix.MatrixHelper[T]
}

func NewGraphHelper[T any](matrixToUse [][]T, equalityChecker func(a, b T) bool) *GraphHelper[T] {
	return &GraphHelper[T]{
		matrixHelper: matrix.NewMatrixHelper(matrixToUse, equalityChecker),
	}
}

// topologicalSortUtil is a recursive helper function for topological sorting.
func (g *GraphHelper[T]) topologicalSortUtil(vertex T, adjacent [][]matrix.Neighbor[T], visited []bool, stack *[]T, columnCount int) {
	pos := g.matrixHelper.GetPositionOfTarget(vertex, nil)
	visited[pos.RowIndex*columnCount+pos.ColIndex] = true

	for _, neighbor := range adjacent[pos.RowIndex*columnCount+pos.ColIndex] {
		if !visited[neighbor.Position.RowIndex*columnCount+neighbor.Position.ColIndex] {
			g.topologicalSortUtil(neighbor.Value, adjacent, visited, stack, columnCount)
		}
	}

	*stack = append(*stack, vertex)
}

// TopologicalSort sorts the vertices of the matrix in a topological order.
func (g *GraphHelper[T]) TopologicalSort() []T {
	adjacent, nodeAmount := g.matrixHelper.GetAdjacencyListHorizontalVertical()

	stack := make([]T, 0)
	visited := make([]bool, nodeAmount)

	columnCount := g.matrixHelper.GetColumnCount()

	for i := 0; i < nodeAmount; i++ {
		if !visited[i] {
			g.topologicalSortUtil(g.matrixHelper.GetAtPosition(i/columnCount, i%columnCount), adjacent, visited, &stack, columnCount)
		}
	}

	for _, v := range stack {
		fmt.Printf("Vertex: %v\n", v)
	}

	fmt.Println(fmt.Sprintf("Number of vertices: %d", len(stack)))

	return stack
}

// Helper function to find the key in a map given the value
func (g *GraphHelper[T]) findKey(m map[int]int, value int) int {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return -1
}

// TODO - Refactor this into PathFinder lib

func (g *GraphHelper[T]) getNeighbors(position matrix.Position, columnCount int, adjacencyList [][]matrix.Neighbor[T]) []matrix.Neighbor[T] {
	rowIndex := position.RowIndex
	colIndex := position.ColIndex
	adjacencyIndex := rowIndex*columnCount + colIndex
	return adjacencyList[adjacencyIndex]
}

func (g *GraphHelper[T]) traverseNodeNeighborsUnique(position matrix.Position, columnCount int, adjacencyList [][]matrix.Neighbor[T], canMoveFunc func(start, end T) bool, currentPath []T, result *[][]T, visited map[matrix.Position]bool) [][]T {
	currentNode := g.matrixHelper.GetAtPosition(position.RowIndex, position.ColIndex)

	if _, ok := visited[position]; ok {
		return *result
	}

	visited[position] = true
	currentPath = append(currentPath, currentNode)

	validNeighbors := false
	for _, neighbor := range g.getNeighbors(position, columnCount, adjacencyList) {
		if canMoveFunc(currentNode, neighbor.Value) {
			validNeighbors = true
			g.traverseNodeNeighborsUnique(neighbor.Position, columnCount, adjacencyList, canMoveFunc, currentPath, result, visited)
		}
	}

	if !validNeighbors && len(currentPath) > 1 {
		*result = append(*result, currentPath)
	}

	return *result
}

func (g *GraphHelper[T]) traverseNodeNeighborsAll(
	position matrix.Position,
	columnCount int,
	adjacencyList [][]matrix.Neighbor[T],
	canMoveFunc func(start, end T) bool,
	currentPath []T,
	allPaths *[][]T,
	visited map[matrix.Position]bool,
	strictComparer func(a, b T) bool,
	endNodeTarget T) [][]T {
	currentNode := g.matrixHelper.GetAtPosition(position.RowIndex, position.ColIndex)

	newPath := make([]T, len(currentPath))
	copy(newPath, currentPath)
	newPath = append(newPath, currentNode)

	for _, neighbor := range g.getNeighbors(position, columnCount, adjacencyList) {
		if canMoveFunc(currentNode, neighbor.Value) {
			g.traverseNodeNeighborsAll(neighbor.Position, columnCount, adjacencyList, canMoveFunc, newPath, allPaths, visited, strictComparer, endNodeTarget)
		}
	}

	lastNode := newPath[len(newPath)-1]
	if strictComparer(lastNode, endNodeTarget) {
		*allPaths = append(*allPaths, newPath)
	}

	return *allPaths
}

func (g *GraphHelper[T]) FindSuitablePathsBetweenNodes(startNode, endNodeTarget T, canMoveFunc func(start, end T) bool, strictComparer func(a, b T) bool, unique bool) [][]T {
	adjacent, _ := g.matrixHelper.GetAdjacencyListHorizontalVertical()

	startNodePositions := g.matrixHelper.GetPositionsOfTarget(startNode, &strictComparer)

	columnCount := g.matrixHelper.GetColumnCount()

	result := make([][]T, 0)

	for _, startPosition := range startNodePositions {
		posResult := make([][]T, 0)
		var validPaths [][]T
		if unique {
			validPaths = g.traverseNodeNeighborsUnique(*startPosition, columnCount, adjacent, canMoveFunc, make([]T, 0), &posResult, make(map[matrix.Position]bool))
		} else {
			validPaths = g.traverseNodeNeighborsAll(*startPosition, columnCount, adjacent, canMoveFunc, make([]T, 0), &posResult, make(map[matrix.Position]bool), strictComparer, endNodeTarget)
		}

		for _, path := range validPaths {
			if strictComparer(path[len(path)-1], endNodeTarget) {
				result = append(result, path)
			}
		}
	}

	return result
}
