package tree_preorder_traversal;

/* you only have to complete the function given below.
Node is defined as

class Node {
    int data;
    Node left;
    Node right;
}

*/

public class Solution {
    public static void preOrder(Node root) {
        if (root == null) return ;

        System.out.printf("%d ", root.data);
        preOrder(root.left);
        preOrder(root.right);
    }

    public static Node insert(Node root, int data) {
        return null; // 適当
    }

    private static class Node {
        int data;
        Node left;
        Node right;
    }
}

