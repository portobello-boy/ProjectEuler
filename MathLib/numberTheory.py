from typing import List, Set
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