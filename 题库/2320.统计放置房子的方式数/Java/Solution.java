/*
 * @lc app=leetcode.cn id=2320 lang=java
 *
 * [2320] 统计放置房子的方式数
 */

// @lc code=start
class Solution {
    static final int MOD = (int) 1e9 + 7, MX = (int) 1e4 + 1;
    static final int[] f = new int[MX];
    
    static {
        f[0] = 1;
        f[1] = 2;
        for (var i = 2; i < MX; ++i)
            f[i] = (f[i - 1] + f[i - 2]) % MOD;
    }

    public int countHousePlacements(int n) {
        return (int) ((long) f[n] * f[n] % MOD);
    }
}
// @lc code=end

