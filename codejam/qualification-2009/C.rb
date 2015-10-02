def gets
  STDIN.gets.strip
end

def analyse(string, pattern)
  todo = {
    "" => 1
  }
  string.each_with_index { |letter, i|
    possibilites = todo.keys.select { |e| pattern[e.size] == letter }
    possibilites.map { |e| e + letter }.each { |k|
      if todo.has_key?(k)
        todo[k] = todo[k] + todo[k[0..-2]]
      else
        todo[k] = todo[k[0..-2]]
      end
    }
    # p todo
  }
  todo[pattern.join]
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "
  puts analyse(gets.split(""), "welcome to code jam".split("")).to_s.rjust(4, "0")[-4..-1]
end
