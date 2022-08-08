import java.util.ArrayList;
import java.util.List;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode(int x) { val = x; }
 * }
 */

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}

class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        if (root == null) {
            return "[]";
        }
        StringBuffer sb = new StringBuffer();
        List<TreeNode> queue = new ArrayList<>();
        queue.add(root);
        int index = 0;
        while (index < queue.size()) {
            TreeNode node = queue.get(index);
            if (node == null) {
                sb.append("#,");
            } else {
                sb.append(node.val + ",");
                queue.add(node.left);
                queue.add(node.right);
            }
            index++;
        }
        return sb.toString();
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        if (data.equals("[]")) {
            return null;
        }
        String[] strs = data.split(",");
        List<TreeNode> queue = new ArrayList<>();
        TreeNode root = new TreeNode(Integer.parseInt(strs[0]));
        queue.add(root);
        int index = 0;
        int qIndex = 0;
        while (qIndex < queue.size()) {
            TreeNode node = queue.get(qIndex);   
            if (!strs[++index].equals("#")) {
                node.left = new TreeNode(Integer.parseInt(strs[index]));
                queue.add(node.left);
            } else {
                node.left = null;
            }
            if (!strs[++index].equals("#")) {
                node.right = new TreeNode(Integer.parseInt(strs[index]));
                queue.add(node.right);
            } else {
                node.right = null;
            }
            qIndex++;
        }
        return root;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));