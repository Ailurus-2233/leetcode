/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}
class Solution {
    int count;
    public int kthLargest(TreeNode root, int k) {
        count = 0;
        return exec(root, k);
    }

    public int exec(TreeNode root, int k) {
        if (root == null) {
            return 0;
        }
        int right = exec(root.right, k);
        if (right != 0) {
            return right;
        }
        count ++;
        if (count == k) {
            return root.val;
        }
        return exec(root.left, k);
    }
}