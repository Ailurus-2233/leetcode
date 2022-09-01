/**
 * Solution
 */
public class Solution {

    public static void main(String[] args) {
        Integer a = 1;
        Integer b = 1;
        Integer c = 500;
        Integer d = 500;
        System.out.println(a == b);
        System.out.println(c == d);
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