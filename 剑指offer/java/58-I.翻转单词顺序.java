import java.util.ArrayList;
import java.util.List;

class Solution {
    public static String reverseWords(String s) {
        String[] arr = s.split(" ");
        List<String> ans = new ArrayList<>();
        for (int i = arr.length-1; i>=0; i--) {
            if (!arr[i].equals("")) {
                ans.add(arr[i]);
            }
        }
        return String.join(" ", ans);
    }

    public static void main(String[] args) {
        System.out.println(reverseWords("a good   example"));
    }
}