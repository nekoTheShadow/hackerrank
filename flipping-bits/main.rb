#!/bin/ruby

require 'json'
require 'stringio'

# Complete the flippingBits function below.
def flippingBits(n)
    n^((1<<32)-1)
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

q = gets.to_i

q.times do |q_itr|
    n = gets.to_i

    result = flippingBits n

    fptr.write result
    fptr.write "\n"
end

fptr.close()

