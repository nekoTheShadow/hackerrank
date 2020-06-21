package delete_a_node_from_a_linked_list;

import java.io.BufferedWriter;
import java.io.IOException;
import java.io.OutputStreamWriter;
import java.util.Scanner;

public class Solution {

    static class SinglyLinkedListNode {
        public int data;
        public SinglyLinkedListNode next;

        public SinglyLinkedListNode(int nodeData) {
            this.data = nodeData;
            this.next = null;
        }
    }

    static class SinglyLinkedList {
        public SinglyLinkedListNode head;
        public SinglyLinkedListNode tail;

        public SinglyLinkedList() {
            this.head = null;
            this.tail = null;
        }

        public void insertNode(int nodeData) {
            SinglyLinkedListNode node = new SinglyLinkedListNode(nodeData);

            if (this.head == null) {
                this.head = node;
            } else {
                this.tail.next = node;
            }

            this.tail = node;
        }
    }

    public static void printSinglyLinkedList(SinglyLinkedListNode node, String sep, BufferedWriter bufferedWriter) throws IOException {
        while (node != null) {
            bufferedWriter.write(String.valueOf(node.data));

            node = node.next;

            if (node != null) {
                bufferedWriter.write(sep);
            }
        }
    }

    // Complete the deleteNode function below.

    /*
     * For your reference:
     *
     * SinglyLinkedListNode {
     *     int data;
     *     SinglyLinkedListNode next;
     * }
     *
     */
    static SinglyLinkedListNode deleteNode(SinglyLinkedListNode head, int position) {
        if (position == 0) {
            return head.next;
        }

        SinglyLinkedListNode node = head;
        for (int i = 0; i < position - 1; i++) {
            node = node.next;
        }
        node.next = node.next.next;

        return head;
    }

    private static final Scanner scanner = new Scanner(System.in);
    public static void main(String[] args) throws IOException {
        int n = scanner.nextInt();
        SinglyLinkedList list = new SinglyLinkedList();
        for (int i = 0; i < n; i++) {
            list.insertNode(scanner.nextInt());
        }
        int position = scanner.nextInt();
        deleteNode(list.head, position);

        BufferedWriter bufferedWriter = new BufferedWriter(new OutputStreamWriter(System.out));

        bufferedWriter.write("Before: ");
        printSinglyLinkedList(list.head, " ", bufferedWriter);
        bufferedWriter.newLine();

        deleteNode(list.head, position);

        bufferedWriter.write("After: ");
        printSinglyLinkedList(list.head, " ", bufferedWriter);
        bufferedWriter.newLine();

        bufferedWriter.flush();
    }
}