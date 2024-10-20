class Solution(object):
    def threeSum(self, nums):
        # Step 1: Sort the array
        nums.sort()
        triplets = []

        # Step 2: Iterate through the array
        for i in range(len(nums) - 2):
            # Step 3: Avoid duplicates for the first number
            if i > 0 and nums[i] == nums[i - 1]:
                continue

            # Step 4: Two-pointer approach
            left, right = i + 1, len(nums) - 1

            while left < right:
                current_sum = nums[i] + nums[left] + nums[right]

                if current_sum == 0:
                    # Found a triplet
                    triplets.append([nums[i], nums[left], nums[right]])

                    # Move the left pointer forward, skipping over duplicates
                    while left < right and nums[left] == nums[left + 1]:
                        left += 1
                    # Move the right pointer backward, skipping over duplicates
                    while left < right and nums[right] == nums[right - 1]:
                        right -= 1

                    # Move both pointers after adding the triplet
                    left += 1
                    right -= 1

                elif current_sum < 0:
                    # If the sum is too small, move the left pointer to the right
                    left += 1
                else:
                    # If the sum is too large, move the right pointer to the left
                    right -= 1

        return triplets
        