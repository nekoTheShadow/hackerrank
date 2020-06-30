# Complete the crosswordPuzzle function below.
def crosswordPuzzle(crossword, words)
  f(crossword, words.strip.split(';'))[0]
end

def f(crossword, words)
  return [crossword] if words.empty?

  ans = []
  word = words[0]
  n = word.size
  [*0..9].product([*0..9]).each do |i, j|
    if i + n <= 10 && n.times.all?{|x| crossword[i+x][j] == '-' || crossword[i+x][j] == word[x]}
      next_crossword = crossword.map(&:dup)
      n.times{|x| next_crossword[i+x][j] = word[x] }
      ans.concat(f(next_crossword, words[1..]))
    end

    if j + n <= 10 && n.times.all?{|y| crossword[i][j+y] == '-' || crossword[i][j+y] == word[y]}
      next_crossword = crossword.map(&:dup)
      n.times{|y| next_crossword[i][j+y] = word[y] }
      ans.concat(f(next_crossword, words[1..]))
    end
  end

  ans
end

crossword = 10.times.map{ gets.strip }
words = gets.strip
p crosswordPuzzle(crossword, words)