def div_while_possible(result, n, prime)
	quot, rem = n.divmod(prime)
	while rem == 0 do
		result.push prime
		quot, rem = quot.divmod(prime)
	end
	return quot * prime + rem
end

n = ARGV.shift.to_i
result = []

n = div_while_possible(result, n, 2)
n = div_while_possible(result, n, 3)

m = 1 # 6m-1, 6m+1のみ考えれば十分
limit = Math::sqrt(n).to_i # nの平方根より大きい素因数は存在しない

while n > limit do
	n = div_while_possible(result, n, 6 * m - 1)
	n = div_while_possible(result, n, 6 * m + 1)

	m = m + 1
end

p result.to_s