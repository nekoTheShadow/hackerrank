# Complete the missingNumbers function below.
def missingNumbers(arr, brr):
    import collections

    d = collections.defaultdict(int)
    for v in brr:
        d[v] += 1
    for v in arr:
        d[v] -= 1

    ans = []
    for k in d:
        if d[k] > 0:
            ans.append(k)

    ans.sort()
    return ans



if __name__ == '__main__':
    n = int(input())
    arr = list(map(int, input().split()))
    m = int(input())
    brr = list(map(int, input().split()))
    print(missingNumbers(arr, brr))