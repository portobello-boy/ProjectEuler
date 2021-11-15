from typing import Dict, Generator, List, Set
from math import floor, sqrt

from .constants import golden, firstPrimes

"""
Sequences and sequence sums
"""

# Return nth triangular number
def triangle(n:int) -> int:
    return int(n * (n+1)/2)

# Return nth Fibonacci numer
def fibonacci(n:int) -> int:
    return int(round((golden ** n) / sqrt(5)))

# Return first n term from Fibonacci Sequence
def fibonacciSequence(n:int) -> List[int]:
    return [fibonacci(i) for i in range(0, n)]

# Return a generator for prime numbers below bound n
def primeSequenceBoundedGenerator(n:int) -> Generator:
    for p in firstPrimes:
        if p > n:
            break

        yield p
    
    k = int((firstPrimes[-1] + 5) / 6)   # Multiple of six (referencing prime 541)
    current = 0                          # Potential prime currently examined

    while current < n:
        current = 6*k - 1

        if current >= n:
            break

        if isPrime(current):
            yield current
        
        current += 2
        if current >= n:
            break

        if isPrime(current):
            yield current
        
        k += 1

# Generate primes below bound n
def primeSequenceBounded(n:int) -> List[int]:
    return list(primeSequenceBoundedGenerator(n))

# Return a generator for n prime numbers
def primeSequenceGenerator(n:int) -> Generator:
    count = 0
    for p in firstPrimes:
        if count > n:
            break
        count += 1
        yield p
    
    k = int((firstPrimes[-1] + 5) / 6)   # Multiple of six (referencing prime 541)
    current = 0                          # Potential prime currently examined

    while count < n:
        current = 6*k - 1

        if isPrime(current):
            count += 1
            yield current
        
        current += 2

        if isPrime(current):
            count += 1
            yield current
        
        k += 1

# Generate n prime numbers
def primeSequence(n:int) -> List[int]:
    return list(primeSequenceGenerator(n))

"""
Tuples
"""

def isPythagoreanTriple(a:int, b:int, c:int) -> bool:
    return a**2 + b**2 == c**2

"""
Divisors and Multiples
"""

# Return a generator for a list of divisors of n
def divisorListGenerator(n:int) -> Generator:
    largerFactors = []
    for i in range(1, int(sqrt(n))):
        if n % i == 0:
            yield i
            if i*i != n:
                largerFactors.append(int(n / i))
    
    for f in reversed(largerFactors):
        yield f

# Return a list of divisors of n
def divisorList(n:int) -> List[int]:
    return list(divisorListGenerator(n))

# Return a map of prime divisors with their powers which divide n
def primeDivisorList(n:int) -> Dict:
    primeFactorMap = {}
    k = 2
    while n != 1:
        while n % k == 0:
            n = n / k
            if not k in primeFactorMap:
                primeFactorMap[k] = 0
            primeFactorMap[k] += 1
        k = nextPrime(k)
    return primeFactorMap

# Return sum of all multiples of n below bound
def sumOfMultsOfN(n:int, bound:int) -> Set[int]:
    return n * triangle(floor(bound/n))

# Return greatest common divisor of u and v
# Using: Binary GCD: https://en.wikipedia.org/wiki/Binary_GCD_algorithm
def gcd(a:int, b:int) -> int:
    if a == b:
        return a
    elif a == 0:
        return b
    elif b == 0:
        return a
    # u is even
    elif a & 1 == 0:
        # v is even
        if b & 1 == 0:
            return 2*gcd(a >> 1, b >> 1)
        # v is odd
        else:
            return gcd(a >> 1, b)
    # u is odd
    elif a & 1 != 0:
        # v is even
        if b & 1 == 0:
            return gcd(a, b >> 1)
        # v is odd and u is greater than v
        elif a > b and b & 1 != 0:
            return gcd((a-b) >> 1, b)
        # v is odd and u is smaller than v
        else:
            return gcd((b-a) >> 1, a)

# Return least common multiple of a and b
def lcm(a:int, b:int) -> int:
    return int((abs(a) / gcd(a, b)) * abs(b))

"""
Primality
"""

# Determine if n is prime
def isPrime(n:int) -> bool:
    if n == 2 or n == 3:
        return True
    if n % 2 == 0 or n % 3 == 0 or n <= 1:
        return False
    for divisor in range(6, min(int(n**0.5) + 6, n-1), 6):
        if n % (divisor - 1) == 0 or n % (divisor + 1) == 0:
            return False
    return True

# Given n, return (n+1)th prime number
def nextPrime(n:int) -> int:
    if n < 0:
        raise ValueError("n must be a non-negative number")
    while True:
        if n % 2 == 0:
            n += 1 # Make n even
        else:
            n += 2
        if isPrime(n):
            return n