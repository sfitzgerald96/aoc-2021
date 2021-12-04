f = File.open('./01/data.txt')

depths = f.read.split("\n").map(&:to_i)
count = 0
# puts depths.count
depths.each_with_index do |_, i|
  if i == 0 
    next
  elsif depths[i] > depths[i-1]
    count += 1
  end
end

puts count
