#!/bin/ruby

require 'json'
require 'stringio'

# Complete the maxSubsetSum function below.
def maxSubsetSum(arr)
  n = arr.size
  return arr[0]  if n == 1
  return arr.max if n == 2
  
  dp = Array.new(n, -Float::INFINITY)
  dp[0] = arr[0]
  dp[1] = [arr[0], arr[1]].max
  (2...n).each do |i|
    dp[i] = [dp[i-1], dp[i], dp[i-2]+arr[i], arr[i]].max
  end
  dp[n-1]
end

n = gets.to_i
arr = gets.strip.split(' ').map(&:to_i)
p maxSubsetSum(arr)

# fptr = File.open(ENV['OUTPUT_PATH'], 'w')

# n = gets.to_i

# arr = gets.rstrip.split(' ').map(&:to_i)

# res = maxSubsetSum arr

# fptr.write res
# fptr.write "\n"

# fptr.close()