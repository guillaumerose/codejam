def gets
  STDIN.gets.strip
end

def bataille(naomi, ken, score)
  if naomi == []
    return score
  else
    carte_naomi, *reste_naomi = *naomi
    cartes_ken_potentielles = ken.select { |carte_ken| carte_ken > carte_naomi }
    if cartes_ken_potentielles == []
      carte_ken = ken.min
      bataille(reste_naomi, ken - [carte_ken], score + 1)
    else
      carte_ken = cartes_ken_potentielles.min
      bataille(reste_naomi, ken - [carte_ken], score)
    end
  end
end

def bataille_menteur(naomi, ken, score)
  if naomi == []
    return score
  else
    carte_ken, *reste_ken = *ken
    if naomi.max > carte_ken
      bataille_menteur(naomi - [naomi.max], reste_ken, score + 1)
    else
      bataille_menteur(naomi - [naomi.min], reste_ken, score)
    end
  end
end

n = gets.to_i
n.times do |i|
  print "Case ##{i + 1}: "
  cards = gets.to_i
  naomi = gets.split(" ").map(&:to_f).sort.reverse
  ken = gets.split(" ").map(&:to_f).sort.reverse
  puts "#{bataille_menteur(naomi, ken, 0)} #{bataille(naomi, ken, 0)}"
end
