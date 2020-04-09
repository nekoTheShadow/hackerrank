package ctci_merge_sort;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Arrays;
import java.util.Scanner;

public class Solution1 {
    // Complete the countInversions function below.
    static long countInversions(int[] arr) {
        long ans = 0;
        int n = Arrays.stream(arr).max().getAsInt();
        BIT bit = new BIT(n);
        for (int i = 0; i < arr.length; i++) {
            ans += bit.sum(n) - bit.sum(arr[i]);
            bit.add(arr[i], 1);
        }
        return ans;
    }

    private static class BIT {
        private int size;
        private int[] tree;

        public BIT(int size) {
            this.size = size;
            this.tree = new int[size + 1];
        }

        public int sum(int i) {
            int s = 0;
            while (i > 0) {
                s += tree[i];
                i -= i & -i;
            }
            return s;
        }

        public void add(int i, int x) {
            while (i <= size) {
                tree[i] += x;
                i += i & -i;
            }
        }
    }

    public static void main(String[] args) throws IOException {
        Scanner stdin = new Scanner(System.in);
        PrintWriter stdout = new PrintWriter(System.out, false);
        int d = stdin.nextInt();
        for (int i = 0; i < d; i++) {
            int n = stdin.nextInt();
            int[] arr = new int[n];
            for (int j = 0; j < n; j++) {
                arr[j] = stdin.nextInt();
            }
            stdout.println(countInversions(arr));
        }
        stdout.flush();
    }
}
