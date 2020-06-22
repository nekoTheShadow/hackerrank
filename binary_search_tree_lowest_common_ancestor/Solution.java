package binary_search_tree_lowest_common_ancestor;

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Scanner;
import java.util.Set;

public class Solution {
    public static Node lca(Node root, int v1, int v2) {
        Map<Integer, Integer> parents = new HashMap<>();
        Map<Integer, Node> nodes = new HashMap<>();
        Deque<Node> q = new ArrayDeque<>();
        parents.put(root.data, -1);
        q.addLast(root);

        while (!q.isEmpty()) {
            Node parent = q.removeFirst();
            nodes.put(parent.data, parent);
            if (parent.left != null) {
                parents.put(parent.left.data, parent.data);
                q.addLast(parent.left);
            }
            if (parent.right != null) {
                parents.put(parent.right.data, parent.data);
                q.addLast(parent.right);
            }
        }

        Set<Integer> ancestors = new HashSet<>();
        while (v1 != -1) {
            ancestors.add(v1);
            v1 = parents.get(v1);
        }
        while (!ancestors.contains(v2)) {
            v2 = parents.get(v2);
        }
        return nodes.get(v2);
    }

    public static Node insert(Node root, int data) {
        if(root == null) {
            return new Node(data);
        } else {
            Node cur;
            if(data <= root.data) {
                cur = insert(root.left, data);
                root.left = cur;
            } else {
                cur = insert(root.right, data);
                root.right = cur;
            }
            return root;
        }
    }

    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int t = scan.nextInt();
        Node root = null;
        while(t-- > 0) {
            int data = scan.nextInt();
            root = insert(root, data);
        }
          int v1 = scan.nextInt();
          int v2 = scan.nextInt();
        scan.close();
        Node ans = lca(root,v1,v2);
        System.out.println(ans.data);
    }
}

class Node {
    Node left;
    Node right;
    int data;

    Node(int data) {
        this.data = data;
        left = null;
        right = null;
    }
}