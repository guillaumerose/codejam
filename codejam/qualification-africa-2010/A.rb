def gets
  STDIN.gets.strip
end

def pick(size)
  a = rand(size)
  b = rand(size)
  while a == b
    b = rand(size)
  end
  [a, b]
end

n = gets.to_i
n.times do |i|
  sum = gets.to_i
  gets # discard
  items = gets.split(" ").map(&:to_i)

  a, b = pick(items.size)
  while items[a] + items[b] != sum
    a, b = pick(items.size)
  end

  print "Case ##{i + 1}: "
  puts a > b ? "#{b + 1} #{a + 1}" : "#{a + 1} #{b + 1}"
end
