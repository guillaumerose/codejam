i = 0
while str = gets
    words = str.strip.split(" ")
    if words.uniq.size == words.size
        i = i + 1
    end
end
puts i