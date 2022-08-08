import java.util.HashSet;

class Solution {
    public boolean isStraight(int[] nums) {
        int max = 0;
        int min = 14;
        int[] count = new int[14];
        for (int num : nums) {
            System.out.println(1);
            if (num > max) {
                max = num;
            }
            if (num < min) {
                min = num;
            }
            count[num]++;
            if (num != 0 && count[num] > 1) {
                return false;
            }
        }
        if (count[0] == 0)
            return (max - min) == 4;
        else
            return (max - min) == 4 - count[0];
    }
}