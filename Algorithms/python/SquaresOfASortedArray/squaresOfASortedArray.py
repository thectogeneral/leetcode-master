def sortedSquares(nums):

    return sorted([a ** 2 for a in nums])
class Solution(object):
    def sortedSquares(self, nums):
        # Map each element 'a' in array 'A' to its square and store the result in a new array
        # Sort the new array in non-decreasing order
        # Return the sorted array
        return sorted([a ** 2 for a in nums])
