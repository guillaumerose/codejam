def gets
  STDIN.gets.strip
end

def scalar_product(a, b)
  a.zip(b).map { |couple| couple[0] * couple[1] }.inject(&:+)
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "

  gets # discard
  x = gets.split(" ").map(&:to_i).sort
  y = gets.split(" ").map(&:to_i).sort.reverse

  puts scalar_product(x, y)
end
