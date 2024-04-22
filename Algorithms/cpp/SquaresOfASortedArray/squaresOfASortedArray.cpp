#include <vector>
#include <algorithm>

class Solution {
public:
    std::vector<int> sortedSquares(std::vector<int>& nums) {
        // Map each element 'num' in vector 'nums' to its square and store the result in a new vector
        // Sort the new vector in non-decreasing order
        // Return the sorted vector
        std::vector<int> squared;
        for (int num : nums) {
            squared.push_back(num * num);
        }
        std::sort(squared.begin(), squared.end());
        return squared;
    }
};
