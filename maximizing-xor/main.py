# Complete the maximizingXor function below.
def maximizingXor(l, r):
    import itertools
    return max(a^b for a, b in itertools.product(range(l, r+1), repeat=2))

if __name__ == "__main__":
    l = int(input())
    r = int(input())
    print(maximizingXor(l, r))
