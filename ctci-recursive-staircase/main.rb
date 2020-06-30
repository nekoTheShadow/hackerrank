@memo = {}

def stepPerms(n)
  return 1 if n == 0
  return 0 if n < 0
  return @memo[n] if @memo.include?(n)
  @memo[n] = (1..3).reduce(0){|sum, x| (sum + stepPerms(n - x)) % (10**10+7) }
end

s = gets.to_i
s.times do 
  n = gets.to_i
  puts(stepPerms(n))
end