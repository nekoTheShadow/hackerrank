# Complete the isBalanced function below.
def isBalanced(s):
    stack = []
    for ch in s:
        print(stack, ch)
        if ch in ('{', '[', '('): 
            stack.append(ch)
        else:
            if len(stack) == 0:
                return 'NO'
            if (stack[-1] == '{' and ch == '}') \
                    or (stack[-1] == '[' and ch == ']') \
                    or (stack[-1] == '(' and ch == ')'):
                stack.pop()
            else:
                return 'NO'
    
    return 'YES' if len(stack) == 0 else 'NO'

if __name__ == '__main__':
    n = int(input())
    for _ in range(n):
        s = input()
        print(isBalanced(s))
