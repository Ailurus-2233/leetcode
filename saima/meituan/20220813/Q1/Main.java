import java.util.Arrays;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int t = sc.nextInt();
        int[] ddl = new int[n];

        for (int i = 0; i < n; i++) {
            ddl[i] = sc.nextInt();
        }

        Arrays.sort(ddl);
        // System.out.println(Arrays.toString(ddl));
        int count = 0;
        for (int i = 0; i < n; i++) {
            if (((i+1-count) * t) <= ddl[i]) {
            } else {
                count++;
            }
        }
        System.out.println(count);
    }

}