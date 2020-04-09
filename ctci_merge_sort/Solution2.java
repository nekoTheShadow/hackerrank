package ctci_merge_sort;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Solution2 {
    // Complete the countInversions function below.
    static long countInversions(int[] arr) {
        List<Integer> a = Arrays.stream(arr).boxed().collect(Collectors.toList());
        return mergeCount(a);
    }

    static long mergeCount(List<Integer> a) {
        int n = a.size();
        if (n <= 1) {
            return 0;
        }

        List<Integer> b = IntStream.range(0, n/2).mapToObj(i -> a.get(i)).collect(Collectors.toList());
        List<Integer> c = IntStream.range(n/2, n).mapToObj(i -> a.get(i)).collect(Collectors.toList());

        long count = 0;
        count += mergeCount(b);
        count += mergeCount(c);

        for (int i = 0, j = 0, k = 0; i < n; i++) {
            if (k == c.size()) {
                a.set(i, b.get(j++));
            } else if (j == b.size()) {
                a.set(i, c.get(k++));
            } else if (b.get(j) <= c.get(k)) {
                a.set(i, b.get(j++));
            } else {
                a.set(i, c.get(k++));
                count += n / 2 - j;
            }
        }

        return count;

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
