def gets
  STDIN.gets.strip
end

def line(row)
  [(row - 1), 0].max.times { line = gets }
  numbers = gets.split(" ").map(&:to_i)
  (4 - row).times { line = gets }
  numbers
end

def configuration
  row = gets.to_i
  first_round = line(row)
  row = gets.to_i
  second_round = line(row)
  first_round & second_round
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "
  output = configuration
  if output.size == 1
    puts output.first
  elsif output.size > 1
    puts "Bad magician!"
  else
    puts "Volunteer cheated!"
  end
end