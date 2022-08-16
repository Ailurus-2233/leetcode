import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Stack;


class Solution {
    public int[] reversePrint(ListNode head) {
        Stack<Integer> stack = new Stack<Integer>();

        List<Integer> a = new ArrayList<>();

        Queue<Integer> q = new LinkedList<>();
        while (head != null) {
            stack.add(head.val);
            head = head.next;
        }
        int[] ans = new int[stack.size()];
        // System.out.println(Arrays.toString(stack.toArray()));
        int size = stack.size();
        for (int i = 0; i < size; i ++) {
            ans[i] = stack.pop();
        }
        return ans;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(2);
        ListNode p = head;
        p.next = new ListNode(3);
        p = p.next;
        p.next = new ListNode(1);

        Solution solution = new Solution();
        
        System.out.println(Arrays.toString(solution.reversePrint(head)));
        
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}