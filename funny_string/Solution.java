package funny_string;

import java.io.*;
import java.util.ArrayList;
import java.util.List;

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader stdin = new BufferedReader(new InputStreamReader(System.in));
        PrintWriter stdout = new PrintWriter(System.out, false);
        new Solution().run(stdin, stdout);
        stdout.flush();
    }

    public void run(BufferedReader stdin, PrintWriter stdout) throws IOException {
        int q = Integer.parseInt(stdin.readLine());
        for (int i = 0; i < q; i++) {
            char[] s = stdin.readLine().toCharArray();

            int n = s.length;
            List<Integer> a1 = new ArrayList<>();
            List<Integer> a2 = new ArrayList<>();
            for (int x = 0; x < n - 1; x++) {
                a1.add(Math.abs(s[x] - s[x+1]));
            }
            for (int x = n - 1; x >= 1; x--) {
                a2.add(Math.abs(s[x] - s[x-1]));
            }
            stdout.println(a1.equals(a2) ? "Funny" : "Not Funny");
        }
    }
}
