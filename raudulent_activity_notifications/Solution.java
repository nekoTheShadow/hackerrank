package raudulent_activity_notifications;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Scanner;

public class Solution {

    // Complete the activityNotifications function below.
    static int activityNotifications(int[] expenditure, int d) {
        List<Integer> a = new ArrayList<>();
        for (int i = 0; i < d; i++) {
            a.add(expenditure[i]);
        }
        Collections.sort(a);

        int count = 0;
        for (int i = d, n = expenditure.length; i < n; i++) {
            if (d % 2 == 0) {
                if (a.get(d / 2 - 1) + a.get(d / 2) <= expenditure[i]) {
                    count++;
                }
            } else {
                if (a.get(d / 2) * 2 <= expenditure[i]) {
                    count++;
                }
            }


            a.remove(Collections.binarySearch(a, expenditure[i - d]));
            int x = Collections.binarySearch(a, expenditure[i]);
            if (x >= 0) {
                a.add(x, expenditure[i]);
            } else {
                a.add(~x, expenditure[i]);
            }
        }

        return count;
    }


    public static void main(String[] args) throws IOException {
        Scanner stdin = new Scanner(System.in);
        int n = stdin.nextInt();
        int d = stdin.nextInt();
        int[] expenditure = new int[n];
        for (int i = 0; i < n; i++) {
            expenditure[i] = stdin.nextInt();
        }
        System.out.println(activityNotifications(expenditure, d));
    }
}
