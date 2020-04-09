n = int(input())
s = input()

ans = 0
h = 0
for c in s:
    if c == 'U':
        h += 1
    else:
        h -= 1
    
    if c == 'D' and h == -1:
        ans += 1

print(ans)