def gets
  STDIN.gets.strip
end

def analyse(pattern, words, index = 0)
  return words if pattern == []
  letter, *tail = *pattern
  if letter == "("
    letters = []
    while letters[-1] != ")"
      letters << tail.shift
    end
    words = words.select { |word| letters[0..-2].include?(word[index]) }
  else
    words = words.select { |word| word[index] == letter }
  end
  analyse(tail, words, index + 1)
end

l, d, n = gets.split(" ").map(&:to_i)
words = (1..d).map { |i| gets.split("") }

n.times do |i|
  print "Case ##{i + 1}: "
  puts analyse(gets.split(""), words).size
end
