function sortedSquares(nums) {
  // Map each element 'a' in array 'A' to its square and store the result in a new array
  // Sort the new array in non-decreasing order
  // Return the sorted array
  return nums.map((a) => a ** 2).sort((a, b) => a - b);
}
