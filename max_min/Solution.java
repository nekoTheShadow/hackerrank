package max_min;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader stdin = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(stdin.readLine());
        int k = Integer.parseInt(stdin.readLine());
        int[] arr = new int[n];
        for (int i= 0; i < n; i++) {
            arr[i] = Integer.parseInt(stdin.readLine());
        }

        Arrays.sort(arr);
        int ret = Integer.MAX_VALUE;
        for (int i = 0; i + k - 1 < n; i++) {
            int min = arr[i];
            int max = arr[i + k - 1];
            ret = Math.min(ret, max - min);
        }
        System.out.println(ret);
    }


}
