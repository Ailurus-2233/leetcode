import java.util.ArrayList;
import java.util.List;

class Solution {
    public int[][] findContinuousSequence(int target) {
        List<int[]> res = new ArrayList<>();
        int left = 1;
        int right = 1;
        int sum = 1;
        while (left <= right) {
            if (sum < target) {
                sum += ++right;
            } else if (sum > target) {
                sum -= left++;
            } else {
                if (left != right) {
                    int[] arr = new int[right - left + 1];
                    for (int i = left; i <= right; i++) {
                        arr[i - left] = i;
                    }
                    res.add(arr);
                }
                sum -= left++;
            }
        }
        return res.toArray(new int[0][]);
    }
}