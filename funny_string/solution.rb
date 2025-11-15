q = gets.to_i
q.times do
    s = gets.chomp

    a1 = s.chars.each_cons(2).map{|x, y| (x.ord - y.ord).abs}
    a2 = s.chars.reverse.each_cons(2).map{|x, y| (x.ord - y.ord).abs}
    puts a1 == a2 ? "Funny" : "Not Funny"
end

