/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/

class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};

class Solution {
    Node pre = null;
    Node first = null;
    boolean flag = true;
    public Node treeToDoublyList(Node root) {
        first = root;
        if (root == null) {
            return null;
        }
        helper(root);
        first.left = pre;
        pre.right = first;
        return first;
    }

    public void helper(Node root) {
        if (root == null) {
            return;
        }
        helper(root.left);
        if (pre != null) {
            if (flag) {
                first = pre;
                flag = false;
            }
            pre.right = root;
            root.left = pre;
        }
        pre = root;
        helper(root.right);
    }
}