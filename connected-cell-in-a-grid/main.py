# Complete the connectedCell function below.
def connectedCell(matrix):
    import itertools
    
    n = len(matrix)
    m = len(matrix[0])
    visited = [[False] * m for _ in range(n)]
    ans = 0

    for x, y in itertools.product(range(n), range(m)):
        if matrix[x][y] == 0 or visited[x][y]:
            continue
            
        visited[x][y] = True
        stack = [(x, y)]
        count = 1
        while stack:
            i, j = stack.pop()
            for di, dj in ((0, 1), (0, -1), (1, 0), (-1, 0), (1, 1), (1, -1), (-1, 1), (-1, -1)):
                ni = i + di
                nj = j + dj
                if 0 <= ni < n and 0 <= nj < m and matrix[ni][nj] == 1 and not visited[ni][nj]:
                    count += 1
                    visited[ni][nj] = True
                    stack.append((ni, nj))
        ans = max(ans, count)
    
    return ans


if __name__ == '__main__':
    n = int(input())
    m = int(input())
    matrix = [list(map(int, input().split())) for _ in range(n)]
    print(connectedCell(matrix))