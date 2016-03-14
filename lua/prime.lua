function is_prime(n)
	local i
	for i = 2, (n-1) do
		if n % i == 0 then return false end
	end
	return true
end

for i=2, 50000 do
	if is_prime(i) then io.write("["..i.."]\n") end
end
