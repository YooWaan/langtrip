class List
  def initialize(arr)
    @arr = arr
  end
  def each
    i = 0
    while i < @arr.length
      yield @arr[i]
      i += 1
    end
  end
end

puts '--- array each ---'
[1, 2, 3].each do |i|
  puts i*2
end

puts '--- class each ---'
list = List.new([4, 5, 6])
list.each do |i|
  puts i*2
end
