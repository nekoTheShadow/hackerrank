#!/bin/ruby

require 'json'
require 'stringio'

# Complete the primality function below.
def primality(n)
    require 'prime'
    n.prime? ? 'Prime' : 'Not prime'
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

p = gets.to_i

p.times do |p_itr|
    n = gets.to_i

    result = primality n

    fptr.write result
    fptr.write "\n"
end

fptr.close()

