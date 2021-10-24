import math

bound = 28123
abundantNums = set()
notAbundantSum = 0

def getPrimeFactors(num):
	factors = {1}
	for i in range(2, int(math.sqrt(num)+1)):
		if num % i == 0:
			factors.add(int(i))
			factors.add(int(num/i))

	return factors

def isAbundant(n):
	return sum(getPrimeFactors(n)) > n

for i in range(12, bound):
    if isAbundant(i):
        abundantNums.add(i)

for i in range(1, bound):
    sumOfAbundants = False
    for abundant in abundantNums:
        if i-abundant in abundantNums:
            sumOfAbundants = True
            break
    if not sumOfAbundants:
        notAbundantSum += i

print(notAbundantSum)