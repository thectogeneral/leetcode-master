func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }

    // Construct adjacency list
    adj := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    // Initialize leaves
    leaves := make([]int, 0)
    for i := 0; i < n; i++ {
        if len(adj[i]) == 1 {
            leaves = append(leaves, i)
        }
    }

    // Keep removing leaves until there are 1 or 2 nodes left
    remainingNodes := n
    for remainingNodes > 2 {
        remainingNodes -= len(leaves)
        newLeaves := make([]int, 0)

        for _, leaf := range leaves {
            neighbor := adj[leaf][0]
            adj[neighbor] = removeElement(adj[neighbor], leaf)
            if len(adj[neighbor]) == 1 {
                newLeaves = append(newLeaves, neighbor)
            }
        }

        leaves = newLeaves
    }

    return leaves
}

func removeElement(slice []int, element int) []int {
    for i := 0; i < len(slice); i++ {
        if slice[i] == element {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}
