package the_love_letter_mystery;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.UncheckedIOException;

public class Solution {
    static void main(String[] args) {
        try {
            new Solution().run();
        } catch (IOException e) {
            throw new UncheckedIOException(e);
        }
    }

    void run() throws IOException {
        BufferedReader stdin = new BufferedReader(new InputStreamReader(System.in));
        int q = Integer.parseInt(stdin.readLine());
        for (int i = 0; i < q; i++) {
            String s = stdin.readLine();
            System.out.println(solve(s));
        }
    }

    int solve(String s) {
        int n = s.length();
        int total = 0;
        for (int i = 0; i < n / 2; i++) {
            char ch1 = s.charAt(i);
            char ch2 = s.charAt(n - i - 1);
            total += Math.abs((int) ch1 - (int) ch2);
        }
        return total;
    }
}
