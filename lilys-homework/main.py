#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the lilysHomework function below.
def lilysHomework(arr):
    c1 = f(arr, True)
    c2 = f(arr, False)
    return min(c1, c2)

def f(arr, reverse):
    n = len(arr)
    a = list(arr)
    b = list(sorted(arr, reverse=reverse))
    d = {a[i] : i for i in range(n)}
    
    c = 0
    for i in range(n):
        if a[i] == b[i]:
            continue
        
        c += 1
        j = d[b[i]]
        a[i], a[j] = a[j], a[i]
        d[a[i]] = i
        d[a[j]] = j
    
    return c



if __name__ == '__main__':
    n = int(input())
    arr = list(map(int, input().rstrip().split()))
    print(lilysHomework(arr))