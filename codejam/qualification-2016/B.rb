def gets
  STDIN.gets.strip
end

def rec(str, i)
  # p str
  # p i
  if str.uniq.size == 1
    return str[0] == "+" ? i : i + 1
  end
  cur = []
  ok = true
  str.each { |e| 
    if ok && e == str[0]
      cur << (e == "-" ? "+" : "-")
    else
      ok = false
      cur << e
    end
  }
  rec(cur, i + 1)
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "
  str = gets.split("")  
  puts rec(str, 0)
end