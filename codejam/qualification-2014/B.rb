def gets
  STDIN.gets.strip
end

def speed(farms, f)
  2 + farms * f
end

def time
  c, f, x = gets.split(" ").map(&:to_f)

  t = 0
  capacite_production = 2
  stock = 0

  while stock < x
    delta_t = ([c, x].min - stock) / capacite_production
    t += delta_t
    stock += delta_t * capacite_production

    if stock < c
      next
    end

    if (x - stock) / capacite_production <= (x - (stock - c)) / (capacite_production + f)
      return t + (x - stock) / capacite_production
    else
      capacite_production += f
      stock -= c
    end
  end
  t
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "
  puts time
end
