from typing import Generator, List, Set
from math import floor, sqrt

from .constants import golden

# Return nth triangular number
def triangle(n:int) -> int:
    return int(n * (n+1)/2)

# Return nth Fibonacci numer
def fibonacci(n:int) -> int:
    return int(round((golden ** n) / sqrt(5)))

# Return first n term from Fibonacci Sequence
def fibonacciSequence(n:int) -> List[int]:
    return [fibonacci(i) for i in range(0, n)]

# Return set of all multiples of n below bound
def multiplesInRange(n:int, bound:int) -> Set[int]:
    return n * triangle(floor(bound/n))

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

# Determine if n is prime
def isPrime(n:int) -> bool:
    if n == 2 or n == 3:
        return True
    if n % 2 == 0 or n % 3 == 0 or n <= 1:
        return False
    for divisor in range(6, min(int(n**0.5) + 6,n-1), 6):
        if n % (divisor - 1) == 0 or n % (divisor + 1) == 0:
            return False
    return True