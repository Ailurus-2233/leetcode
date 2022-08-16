package 题库;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

/*
 * @lc app=leetcode.cn id=15 lang=java
 *
 * [15] 三数之和
 */

// @lc code=start
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        if (nums.length < 3) {
            return null;
        }
        Arrays.sort(nums);
        Map<String, List<Integer>> ans = new HashMap<>();
        Integer[] tmp = null;
        for (int i = 0; i < nums.length; i++) {
            int left = i + 1;
            int right = nums.length-1;
            while(left < right) {
                int sum = nums[i] + nums[left] + nums[right];
                if (sum == 0) {
                    tmp = new Integer[]{nums[i],nums[left],nums[right]};
                    ans.put(Arrays.toString(tmp), Arrays.asList(tmp));
                    right--;
                    left++;
                } else if (sum > 0) {
                    right--;
                } else {
                    left++;
                }
                
            }
        }
        return new ArrayList<List<Integer>>(ans.values());
    }
}
// @lc code=end

