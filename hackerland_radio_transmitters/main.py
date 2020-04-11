def hackerlandRadioTransmitters(x, k):
    import bisect
    
    x.sort()
    
    p = 0
    ans = 0
    while True:
        ans += 1

        q = bisect.bisect_left(x, x[p] + k)
        if q == len(x):
            break

        if x[q] != x[p] + k:
            q -= 1

        while p < len(x):
            if x[q] + k < x[p]:
                break
            p += 1

        if len(x) <= p:
            break

    return ans


if __name__ == '__main__':
    n, k = map(int, input().split())
    x = list(map(int, input().split()))
    print(hackerlandRadioTransmitters(x, k))