def isValid(s):
    import collections

    c = collections.Counter(s)
    v = list(sorted(c.values()))

    # すべての文字種の登場回数がすべて同じ場合
    if len(set(v)) == 1:
        return "YES"
    
    # 最小値が1で、残りすべてが同じ場合
    if v[0] == 1 and len(set(v[1:])) == 1:
        return "YES"

    # 最大値がxで、残りすべてがx-1の場合
    if all(v[i] == v[-1] - 1 for i in range(0, len(v)-1)):
        return "YES"
    
    return "NO"


if __name__ == '__main__':
    s = input()
    print(isValid(s))