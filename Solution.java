/**
 * Solution
 */
public class Solution {

    public static void main(String[] args) {
        System.out.println(func("abs"));
    }

    public static String func(String s) {
        char[] chars = new char[s.length()];
        int index = s.length();
        for (char c: s.toCharArray()) {
            chars[--index] = c;
        }
        return String.valueOf(chars);
    }
}