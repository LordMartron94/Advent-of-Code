package transforming

// GraphHelper is a helper struct for working with a graph data structure.
// Inspired by https://www.geeksforgeeks.org/topological-sorting/
type GraphHelper struct {
}

func NewGraphHelper() *GraphHelper {
	return &GraphHelper{}
}

// TopologicalSortUtil is a recursive helper function for topological sorting.
func (g *GraphHelper) TopologicalSortUtil(vertex int, adjacent [][]int, visited []bool, stack *[]int) {
	visited[vertex] = true

	for _, neighbor := range adjacent[vertex] {
		if !visited[neighbor] {
			g.TopologicalSortUtil(neighbor, adjacent, visited, stack)
		}
	}

	*stack = append(*stack, vertex)
}

// TopologicalSort sorts the vertices of a given graph in a topological order.
func (g *GraphHelper) TopologicalSort(adjacent [][]int, nodeNumber int) []int {
	stack := make([]int, 0)
	visited := make([]bool, nodeNumber)

	for i := 0; i < nodeNumber; i++ {
		if !visited[i] {
			g.TopologicalSortUtil(i, adjacent, visited, &stack)
		}
	}

	return stack
}

func (g *GraphHelper) TopologicalSortPairs(pairs [][]int) []int {
	// 1. Create a map to store nodes and their indices
	nodeMap := make(map[int]int)
	nodeCount := 0
	for _, pair := range pairs {
		for _, node := range pair {
			if _, exists := nodeMap[node]; !exists {
				nodeMap[node] = nodeCount
				nodeCount++
			}
		}
	}

	// 2. Create the adjacency list
	adjacent := make([][]int, nodeCount)
	for i := range adjacent {
		adjacent[i] = make([]int, 0)
	}
	for _, pair := range pairs {
		src := nodeMap[pair[0]]
		dest := nodeMap[pair[1]]
		adjacent[src] = append(adjacent[src], dest)
	}

	// 3. Call TopologicalSort with the generated adjacency list
	stack := make([]int, 0)
	visited := make([]bool, nodeCount)

	for i := 0; i < nodeCount; i++ {
		if !visited[i] {
			g.TopologicalSortUtil(i, adjacent, visited, &stack)
		}
	}

	// 4. Map the indices back to original node values
	result := make([]int, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, g.findKey(nodeMap, stack[i]))
	}

	return result
}

// Helper function to find the key in a map given the value
func (g *GraphHelper) findKey(m map[int]int, value int) int {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return -1
}
