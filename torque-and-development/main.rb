# Complete the roadsAndLibraries function below.
def roadsAndLibraries(n, c_lib, c_road, cities)
  return n*c_lib if c_lib <= c_road

  nodes = Array.new(n){ Array.new }
  cities.each do |(u, v)|
    nodes[u-1] << v-1
    nodes[v-1] << u-1
  end

  visited = Array.new(n, false)
  cost = 0
  n.times do |i|
    next if visited[i]

    cost += c_lib

    queue = [i]
    visited[i] = true
    until queue.empty?
      j = queue.shift

      visited[j] = true
      nodes[j].each do |k|
        next if visited[k]
        cost += c_road
        visited[k] = true
        queue << k
      end
    end
  end

  cost
end

require "stringio"

def execute()
  q = gets.to_i
  q.times do 
    n, m, c_lib, c_road = gets.split(' ').map(&:to_i)
    cities = m.times.map{ gets.split(' ').map(&:to_i) }
    p roadsAndLibraries(n, c_lib, c_road, cities)
  end
end

if __FILE__ == $0
  buffer = StringIO.new
  $stdout = buffer
  Signal.trap(:EXIT) do 
    $stdout = STDOUT
    buffer.rewind
    buffer.each{|line| puts line }
  end
  execute
end





