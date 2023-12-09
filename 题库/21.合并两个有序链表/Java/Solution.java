package 题库;
/*
 * @lc app=leetcode.cn id=21 lang=java
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode() {}
 * ListNode(int val) { this.val = val; }
 * ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        // ListNode zero = new ListNode(101);
        // ListNode head = new ListNode();
        // ListNode p = head;
        // if (list1 == null) list1 = zero;
        // if (list2 == null) list2 = zero;
        // while(list1 != zero || list2 != zero) {
        //     if (list1.val < list2.val) {
        //         p.next = list1;
        //         list1 = list1.next;
        //     } else {
        //         p.next = list2;
        //         list2 = list2.next;
        //     }
        //     p = p.next;
        //     p.next = null;
        //     if (list1 == null) list1 = zero;
        //     if (list2 == null) list2 = zero;
        // }
        // return head.next;
        ListNode head = new ListNode();
        ListNode p = head;
        while (list1 != null && list2 != null) {
            if (list1.val < list2.val) {
                p.next = list1;
                list1 = list1.next;
            } else {
                p.next = list2;
                list2 = list2.next;
            }
            p = p.next;
            p.next = null;
        }
        if (list1 == null) {
            p.next = list2;
        } else {
            p.next = list1;
        }
        return head.next;
    }
}
// @lc code=end

class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}
