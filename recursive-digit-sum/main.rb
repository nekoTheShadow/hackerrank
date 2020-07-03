#!/bin/ruby

# 参考
# https://ja.wikipedia.org/wiki/%E6%95%B0%E5%AD%97%E6%A0%B9

def superDigit(n, k)
  n = n.to_i
  return 0 if n == 0
  m = n * k
  (m % 9 == 0) ? 9 : m % 9
end

tokens = gets.split(' ')
n = tokens[0]
k = tokens[1].to_i
p superDigit(n, k)

