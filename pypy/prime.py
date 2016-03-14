
def is_prime(n):
	for i in range(2, n-1):
		if n % i == 0 :
			return False;
	return True;

primes = []
for i in range(2,50001):
	if is_prime(i) :
		print("[%d]" % i)
		#primes.append(i)
