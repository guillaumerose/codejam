require 'set'

i = 0
while str = gets
    words = str.strip.split(" ")
    if words.uniq.size == words.size
        seen = {}
        words.each do |w|
            seen[Set.new(w.split(''))] = true
        end
        if seen.size == words.size
            i = i + 1
        end
    end
end
puts i