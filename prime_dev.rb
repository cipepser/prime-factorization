n = ARGV.shift.to_i
result = []
prime = 2

while n != 1 do
	quot, rem = n.divmod(prime)
	while rem == 0 do
		result.push prime
		n = quot
		quot, rem = n.divmod(prime)
	end
	prime = prime + 1
end

p result.to_s