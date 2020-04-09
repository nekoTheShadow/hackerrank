import collections

c = collections.Counter()
d = collections.Counter()

q = int(input())
for _ in range(q):
    query, x = map(int, input().split())
    
    if query == 1:
        c[x] += 1
        d[c[x]] += 1
        if d[c[x] - 1] > 0:
            d[c[x] - 1] -= 1
    
    if query == 2:
        if c[x] > 0:
            c[x] -= 1
            d[c[x]] += 1
            d[c[x] + 1] -= 1
            
    
    if query == 3:
        if d[x] > 0:
            print(1)
        else:
            print(0)