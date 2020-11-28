import math

num = 10000

def getPrimeFactors(num):
	factors = {1}
	for i in range(2, int(math.sqrt(num)+1)):
		if num % i == 0:
			factors.add(int(i))
			factors.add(int(num/i))

	return factors

def d(n):
	return sum(getPrimeFactors(n))

amicableNums = {0}
amicableSum = 0
for a in range(2, num + 1):
	if d(d(a)) == a and a != d(a):
		# print(a, d(a))
		amicableNums.add(a)
		amicableNums.add(d(d(a)))

print(sorted(amicableNums))
print(sum(amicableNums))