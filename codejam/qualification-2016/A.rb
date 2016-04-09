def gets
  STDIN.gets.strip
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "
  n = gets.to_i
  acc = []
  j = 1
  if n == 0
    puts "INSOMNIA"
    next
  end
  while true
    if acc.size == 10
      puts n * (j - 1) 
      break
    end
    acc.push(*(j * n).to_s.split("").uniq)
    acc.uniq!
    j += 1
  end
end