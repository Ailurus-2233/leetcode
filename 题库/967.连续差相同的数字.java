import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=967 lang=java
 *
 * [967] 连续差相同的数字
 */

// @lc code=start
class Solution {
    public int[] numsSameConsecDiff(int n, int k) {
        if (n == 1)
            return new int[] { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 };
        var ans = new ArrayList<Integer>();
        for (var i = 1; i < 10; ++i)
            dfs(i, n - 1, k, ans);
        return ans.stream().mapToInt(Integer::intValue).toArray();
    }

    public void dfs(int num, int n, int k, ArrayList<Integer> ans) {
        if (n == 0) {
            ans.add(num);
            return;
        }
        var last = num % 10;
        if (last + k < 10)
            dfs(num * 10 + last + k, n - 1, k, ans);
        if (k != 0 && last - k >= 0)
            dfs(num * 10 + last - k, n - 1, k, ans);
    }
}
// @lc code=end

