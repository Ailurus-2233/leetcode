/*
 * @lc app=leetcode.cn id=1374 lang=java
 *
 * [1374] 生成每种字符都是奇数个的字符串
 */

// @lc code=start
class Solution {
    public String generateTheString(int n) {
        StringBuilder sb = new StringBuilder();
        if (n % 2 == 0) {
            sb.append('a');
            for (int i = 1; i < n; i++) {
                sb.append('b');
            }
        } else {
            for (int i = 0; i < n - 2; i++) {
                sb.append('a');
            }
            if (n > 1) {
                sb.append('b');
            }
            sb.append('c');
        }
        return sb.toString();
    }
}
// @lc code=end

