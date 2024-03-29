class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        d = dict()
        for i, n in enumerate(nums):
            if target - n in d:
                return [d[target - n], i]
            d[n] = i
        return []
