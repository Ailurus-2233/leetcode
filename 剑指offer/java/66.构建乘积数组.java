class Solution {
    public int[] constructArr(int[] a) {
        if (a.length == 0)
            return new int[0];
        int[] ans = new int[a.length];
        ans[0] = 1;
        int tmp = 1;
        for (int i = 1; i < a.length; i++) {
            ans[i] = a[i - 1] * ans[i - 1];
        }
        for (int i = a.length - 2; i >= 0; i++) {
            tmp *= a[i+1];
            ans[i] *= tmp;
        }
        return ans;
    }
}