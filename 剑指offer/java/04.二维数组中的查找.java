class Solution {
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        for (int i = matrix.length-1; i >=0; i--) {
            for (int j = 0; j < matrix[0].length; j++) {
                if (matrix[i][j] == target) {
                    return true;
                } else {
                    if (matrix[i][j] > target) {
                        break;
                    }
                }
            }
        }
        return false;
    }
}