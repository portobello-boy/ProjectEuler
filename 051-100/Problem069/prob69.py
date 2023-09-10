from MathLib.numberTheory import primeSieve
from math import prod

BOUND = 1_000_000

primes = primeSieve(BOUND)

print(primes)
        
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

def totient(a:int) -> int:
    return round(a * prod([(1.0 - 1.0/p) for p in primeDivisorList(a).keys()]))

maxN = 2
maxFunc = float(maxN)/totient(maxN)

for i in range(2, BOUND+1):
    tot = totient(i)
    if i % 1000 == 0:
        print(i, tot)
    if float(i)/float(tot) > maxFunc:
        maxN, maxFunc = i, float(i)/float(tot)
        
print(maxN, maxFunc)