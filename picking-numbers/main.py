#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'pickingNumbers' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY a as parameter.
#

def pickingNumbers(a):
    import collections

    c = collections.Counter(a)
    keys = list(sorted(c))
    ans = max(c[k] for k in keys)
    for i in range(len(keys)-1):
        k1 = keys[i]
        k2 = keys[i+1]
        if k1 + 1 == k2:
            ans = max(ans, c[k1] + c[k2])
    return ans

if __name__ == '__main__':
    n = int(input())
    a = list(map(int, input().split()))
    print(pickingNumbers(a))