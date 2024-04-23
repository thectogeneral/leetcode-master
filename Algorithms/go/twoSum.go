func twoSum(nums []int, target int) []int {
    // Create a map to store the complement of each number
    complementMap := make(map[int]int)

    // Iterate through the nums array
    for i, num := range nums {
        // Calculate the complement
        complement := target - num

        // Check if the complement exists in the map
        if index, found := complementMap[complement]; found {
            // If found, return the indices
            return []int{index, i}
        }

        // Otherwise, add the current number and its index to the map
        complementMap[num] = i
    }

    // Return an empty slice if no solution is found
    return nil
}
