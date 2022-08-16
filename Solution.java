import java.util.Arrays;
import java.util.HashSet;

public class Solution {

    public static void main(String[] args) {
        Integer[] a = {1, 2, 2, 3, 4};
        HashSet<Integer> hs = new HashSet<Integer>(Arrays.asList(a));
        Object[] b =  hs.toArray();
        System.out.println(Arrays.toString(b));
    }
}
