import java.util.Deque;
import java.util.LinkedList;

class Solution {
    public static int[] maxSlidingWindow(int[] nums, int k) {
        if (nums.length == 0 || k == 0)
            return new int[0];
        Deque<Integer> deque = new LinkedList<>();
        int[] res = new int[nums.length - k + 1];
        for (int i = 0, j = 1 - k; i < nums.length; i++, j++) {
            // 如果是队首则推出队列
            if (j >= 0 && nums[j] == deque.peekFirst()) {
                deque.removeFirst();
            }
            // 判断当前元素添加到队列时是否为除队首元素最大的元素
            while (!deque.isEmpty() && deque.peekLast() < nums[i]) {
                deque.removeLast();
            }
            // 添加新元素
            deque.add(nums[i]);
            // 添加队首元素到res
            if (j >= 0) {
                res[j] = deque.peekFirst();
            }
        }
        return res;
    }

    public static void main(String[] args) {
        maxSlidingWindow(new int[] { 1, 3, -1, -3, -2, 4, 2, 1, 2, 5, 3, 6, 7 }, 3);
    }
}