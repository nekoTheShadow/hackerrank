#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the balancedSums function below.
def balancedSums(arr):
    left = 0
    right = sum(arr)
    for i in range(len(arr)):
        if left == right - arr[i]:
            return "YES"
        left += arr[i]
        right -= arr[i]
    return "NO"

if __name__ == '__main__':
    t = int(input())
    for _ in range(t):
        n = int(input())
        arr = list(map(int, input().split()))
