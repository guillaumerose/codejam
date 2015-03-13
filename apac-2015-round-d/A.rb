require 'matrix'
require 'ap'

def gets
  STDIN.gets.strip
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "
  mat = []
  gets.to_i.times do |j|
    mat << gets.split(" ").map(&:to_i)
  end
  connections = []
  mat.each do |line|
    line.each_cons(2) do |x, y|
      if (x - y).abs == 1
        connections << [x, y].sort
      end
    end
  end
  Matrix[*mat].transpose.to_a.each do |line|
    line.each_cons(2) do |x, y|
      if (x - y).abs == 1
        connections << [x, y].sort
      end
    end
  end
  hash = {}
  connections.sort_by(&:first).each do |conn|
    link = hash.find { |k, v| v == conn[0]}
    if link != nil
      hash[link[0]] = conn[1]
    else
      hash[conn[0]] = conn[1]
    end
  end
  a = hash.map { |k, v| [v - k, -k] }.sort
  if a.last
    print "#{-1*a.last[1]} #{a.last[0] + 1}"
  else
    print "1 1"
  end
  puts
end