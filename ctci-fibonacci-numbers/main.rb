@memo = {}

def fibonacci(n)
  return 0 if n == 0
  return 1 if n == 1
  @memo[n] ||= fibonacci(n - 2) + fibonacci(n - 1)
end

n = gets.to_i
print(fibonacci(n))