package maximum_xor;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.io.UncheckedIOException;
import java.lang.reflect.Array;
import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;
import java.util.Objects;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

public class Solution {
    // Complete the maxXor function below.
    static int[] maxXor(int[] arr, int[] queries) {
        Node root = new Node();

        for (int v : arr) {
            Node node = root;
            for (int i = 30; i >= 0; i--) {
                if ((v & (1<<i)) == 0) { // bit=0
                    if (node.left == null) {
                        node.left = new Node();
                    }
                    node = node.left;
                } else {
                    if (node.right == null) {
                        node.right = new Node();
                    }
                    node = node.right;
                }
            }
            node.digit = v;
        }

        int n = queries.length;
        int[] ans = new int[n];
        for (int x = 0; x < n; x++) {
            int v = queries[x];
            Node node = root;
            for (int i = 30; i >= 0; i--) {
                if ((v & (1<<i)) == 0) { // bit=0
                    if (node.right == null) {
                         node = node.left;
                    } else {
                         node = node.right;
                    }
                } else {
                    if (node.left == null) {
                        node = node.right;
                   } else {
                        node = node.left;
                   }
                }
            }

            ans[x] = node.digit ^ v;
        }

        return ans;
    }

    private static class Node {
        private Node left;
        private Node right;
        private int digit;
    }

    public void exec(Stdin stdin, Stdout stdout) {
        int n = stdin.nextInt();
        int[] arr = stdin.nextIntArray(n);
        int m = stdin.nextInt();
        int[] queries = stdin.nextIntArray(m);
        stdout.println(maxXor(arr, queries));
    }

    public static void main(String[] args) {
        Stdin stdin = new Stdin();
        Stdout stdout = new Stdout();
        new Solution().exec(stdin, stdout);
        stdout.flush();
    }

    public static class Stdin {
        private BufferedReader stdin;
        private Deque<String> tokens;
        private Pattern delim;

        public Stdin() {
            stdin = new BufferedReader(new InputStreamReader(System.in));
            tokens = new ArrayDeque<>();
            delim = Pattern.compile(" ");
        }

        public String nextString() {
            try {
                if (tokens.isEmpty()) {
                    String line = stdin.readLine();
                    delim.splitAsStream(line).forEach(tokens::addLast);
                }
                return tokens.pollFirst();
            } catch (IOException e) {
                throw new UncheckedIOException(e);
            }
        }

        public int nextInt() {
            return Integer.parseInt(nextString());
        }

        public double nextDouble() {
            return Double.parseDouble(nextString());
        }

        public int[] nextIntArray(int n) {
            int[] a = new int[n];
            for (int i = 0; i < n; i++) {
                a[i] = nextInt();
            }
            return a;
        }

        public double[] nextDoubleArray(int n) {
            double[] a = new double[n];
            for (int i = 0; i < n; i++) {
                a[i] = nextDouble();
            }
            return a;
        }
    }

    public static class Stdout {
        private PrintWriter stdout;

        public Stdout() {
            stdout =  new PrintWriter(System.out, false);
        }

        public void printf(String format, Object ... args) {
            stdout.printf(format, args);
        }

        public void println(Object ... objs) {
            String line = Arrays.stream(objs).map(this::deepToString).collect(Collectors.joining(" "));
            stdout.println(line);
        }

        private String deepToString(Object o) {
            if (o == null || !o.getClass().isArray()) {
                return Objects.toString(o);
            }

            int len = Array.getLength(o);
            String[] tokens = new String[len];
            for (int i = 0; i < len; i++) {
                tokens[i] = deepToString(Array.get(o, i));
            }
            return "{" + String.join(",", tokens) + "}";
        }

        public void flush() {
            stdout.flush();
        }
    }
}
