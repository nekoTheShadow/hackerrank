package insert_a_node_at_a_specific_position_in_a_linked_list;

import java.io.BufferedWriter;
import java.io.IOException;
import java.io.PrintWriter;
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

    // Complete the insertNodeAtPosition function below.

    /*
     * For your reference:
     *
     * SinglyLinkedListNode {
     *     int data;
     *     SinglyLinkedListNode next;
     * }
     *
     */
    static SinglyLinkedListNode insertNodeAtPosition(SinglyLinkedListNode head, int data, int position) {
        if (position == 0) {
            SinglyLinkedListNode node = new SinglyLinkedListNode(data);
            node.next = head;
            return node;
        }

        SinglyLinkedListNode node1 = head;
        for (int i = 0; i < position-1; i++) {
            node1 = node1.next;
        }
        SinglyLinkedListNode node2 = new SinglyLinkedListNode(data);
        node2.next = node1.next;
        node1.next = node2;
        return head;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        int n = scanner.nextInt();
        SinglyLinkedList list = new SinglyLinkedList();
        for (int i = 0; i < n; i++) {
            list.insertNode(scanner.nextInt());
        }

        int data = scanner.nextInt();
        int position = scanner.nextInt();
        insertNodeAtPosition(list.head, data, position);

        try (BufferedWriter stdout = new BufferedWriter(new PrintWriter(System.out))) {
            printSinglyLinkedList(list.head, " ", stdout);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}