require 'prime'

def gets
  STDIN.gets.strip
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "
  n, j = gets.split(" ").map { |e| e.to_i }
  
  r = {}
  checked = {}
  while r.size != j
    puts r.size
    current = ([1] + (2..n - 1).map { |e| rand(0..1) } + [1]).join
    next if checked.has_key?(current)
    checked[current] = 1
    if !r.has_key?(current) && (2..10).map { |e| Prime.prime?(current.to_i(e)) }.uniq == [false]
      r[current] = "#{current} #{(2..10).map { |e| current.to_i(e).prime_division.first.first }.join(" ")}"
    end
  end
  r.values.each { |e| puts e }
end