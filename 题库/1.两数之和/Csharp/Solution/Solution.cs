using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Solution
{
    public class Solution
    {
        public int[] TwoSum(int[] nums, int target)
        {
            var dict = new Dictionary<int, int>();
            for (var i = 0; i < nums.Length; i++)
            {
                var num = nums[i];
                var diff = target - num;
                if (dict.TryGetValue(diff, out int value))
                {
                    return [value, i];
                }
                dict[num] = i;
            }
            return [-1, -1];
        }
    }
}