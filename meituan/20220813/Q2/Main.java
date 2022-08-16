import java.util.Arrays;
import java.util.Scanner;

class LinkedNode {
    int val;
    LinkedNode next;
    LinkedNode pre;
    LinkedNode(int val) {
        this.val = val;
    }
}

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] s = new int[n];
        for (int i = 0; i < n; i++) {
            s[i] = sc.nextInt();
        }
        LinkedNode tail = new LinkedNode(s[n-1]);
        LinkedNode head = new LinkedNode(s[n-2]);
        LinkedNode cut = head;
        head.next = tail;
        tail.pre = head;

        for (int i = n-3; i >= 0; i--) {
            head.pre = new LinkedNode(s[i]);
            head.pre.next = head;
            head = head.pre;
            LinkedNode tmp = cut.pre;
            head = move(head, cut, tail);
            cut = tmp.pre;
            tail = tmp;
        }
        print(head);

    }

    public static LinkedNode move(LinkedNode head, LinkedNode cut, LinkedNode tail) {
        tail.next = head;
        head.pre = tail;
        head = cut;
        tail = cut.pre;
        head.pre = null;
        tail.next = null;
        return head;
    }

    public static void print(LinkedNode ln) {
        while (ln != null) {
            System.out.print(ln.val + " ");
            ln = ln.next;
        }
        System.out.println();
    }
}
