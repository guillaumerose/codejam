require 'ap'

def gets
  STDIN.gets.strip
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}:"
  gets
  buses = gets.split(" ").map(&:to_i).each_slice(2).to_a
  gets.to_i.times { |p| 
    question = gets.to_i
    print " "
    print buses.select { |b| b[0] <= question && question <= b[1] }.size
  }
  puts
  gets
end