from MathLib.numberTheory import primeSieve
from MathLib.digits import isPermutation
from math import prod

BOUND = 10_000_000

primes = primeSieve(BOUND)

# print(primes)
        
def primeDivisorList(n:int):
    """
    Return a map of prime divisors with their powers which divide n

    Arguments:
    n -- integer for which to return the prime divisors
    """
    primeFactorMap = {}
    
    # if n in primes:
    #     return {n: 1}
    
    for p in primes:
        if n % p == 0:
            while n % p == 0:
                n = n / p # type: ignore
                if not p in primeFactorMap:
                    primeFactorMap[p] = 0
                primeFactorMap[p] += 1
        if n == 1:
            return primeFactorMap       
    return primeFactorMap

exploredTotients = {}
def totient(a:int) -> int:
    if a > 4 and a % 4 == 0:
        totient = 2*exploredTotients[a//2] # type: ignore
    elif a > 4 and a % 2 == 0:
        totient = exploredTotients[a//2] # type: ignore
    else:
        totient = round(a * prod([(1.0 - 1.0/p) for p in primeDivisorList(a).keys()]))
    exploredTotients[a] = totient
    return totient

minN = 2
minFunc = float(minN)/totient(minN)

for i in range(2, BOUND+1):
    tot = totient(i)
    if not isPermutation(i, tot):
        continue
    print(i, tot)
    if float(i)/float(tot) < minFunc:
        minN, minFunc = i, float(i)/float(tot)
        
print(minN, minFunc)