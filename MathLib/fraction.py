from typing import Optional, Sequence
from math import sqrt, floor
from MathLib.numberTheory import gcd
from decimal import Decimal, getcontext

getcontext().prec = 100

class Fraction:
    """
    Structure which contains a fraction and methods to be evaluated on it
    """
    def __init__(self, numerator, denominator = 1):
        self.n = numerator
        self.d = denominator
    
    def __add__(self, other:'Fraction') -> 'Fraction':
        greatestCommonDivisor = gcd(self.d, other.d)

        n1 = self.n * (greatestCommonDivisor // self.d)
        n2 = other.n * (greatestCommonDivisor // other.d)
        
        return Fraction(n1 + n2, greatestCommonDivisor)

    def __mul__(self, other:'Fraction') -> 'Fraction':
        return Fraction(self.n * other.n, self.d * other.d)

    def addInt(self, n:int) -> 'Fraction':
        return Fraction(self.n + (n * self.d), self.d)

    def reciprocal(self):
        return Fraction(self.d, self.n)

    def str(self):
        return "{}/{}".format(self.n, self.d)

class ContinuedFraction:
    """
    Structure which contains a continued fraction and methods to be evaluated on it
    """
    def __init__(self, addend:int = 0,
                 numerator:int = 0, denominator:Optional['ContinuedFraction']=None,
                 infinite:bool = False):
        self.a = addend
        self.n = numerator
        self.d = denominator
        self.r = 0
        self.inf = infinite
        
    def __init__(self, number:float = 0, depth:int = 10):
        epsilon = 10**-30
        self.inf = False
        self.a = int(floor(number))
        r = number - floor(number)
        self.r = r
        if r < epsilon:
            self.n = 0
            self.r = 0
            self.d = None
        else:
            self.n = 1
            if depth == 1:
                self.d = None
            else:
                self.d = ContinuedFraction(Decimal('1.0')/Decimal(str(r)), depth=depth-1)

    def copy(self) -> 'ContinuedFraction':
        """
        Returns a shallow copy of self
        """
        return ContinuedFraction(addend = self.a, numerator = self.n,
                                 denominator = self.d, infinite = self.inf)

    def evaluate(self, depth:int = 10) -> float:
        """
        Recursively evaluate the continued fraction.
        If the fraction is infinite, recursively evaluate using the previous addend.

        Arguments:
        depth -- depth to recurse to if fraction is infinite
        """
        # print("a: {}, n: {}, d: {}, depth: {}".format(self.a, self.n, self.d, depth))

        if not depth == 0 and self.inf:
            # print("INFINITE")
            return self.a + (self.n / self.evaluate(depth = depth-1))

        if (depth == 0 and self.inf) or self.n == 0 or self.d is None:
            return float(self.a)

        return self.a + (self.n / self.d.evaluate(depth = depth))

    def evaluateFraction(self, depth:int = 10) -> 'Fraction':
        """
        Recursively evaluate the continued fraction and return a fraction.
        If the fraction is infinite, recursively evaluate using the previous addend.

        Arguments:
        depth -- depth to recurse to if fraction is infinite
        """
        # print("a: {}, n: {}, d: {}, inf: {}, depth: {}".format(self.a, self.n, self.d, self.inf, depth))

        if depth == 0 or self.d == None:
            return Fraction(self.a)

        if not depth == 0 and self.inf:
            # print("INFINITE")
            subFrac = Fraction(self.n) * self.evaluateFraction(depth = depth-1).reciprocal()
            return subFrac.addInt(self.a)
        
        if (depth == 0 and self.inf) or self.n == 0 or self.d is None:
            return Fraction(self.a)
            
        subFrac = Fraction(self.n) * self.d.evaluateFraction(depth = depth-1).reciprocal()
        return subFrac.addInt(self.a)

    @staticmethod
    def generateFromSequence(sequence:Sequence[int], numerator:int=1, infinite:bool=False):
        """
        Generates a continued fraction from a sequence of integers.

        Arguments:
        sequence -- list of integers with addends
        numerator -- numerator for each continued fraction
        infinite -- flag to determine if fraction is infinite
        """

        if len(sequence) == 0:
            return ContinuedFraction()

        fraction = ContinuedFraction(addend=sequence[-1], infinite=infinite)

        for num in reversed(sequence[:-1]):
            if not isinstance(num, int):
                raise TypeError
            fraction.d = fraction.copy()
            fraction.n = numerator
            fraction.a = num
            fraction.inf = False # Ensure that the infinite fraction is only at end
        
        return fraction

    def getSequence(self) -> Sequence[int]:
        """
        Return sequence of integers which constiture this continued fraction
        """

        sequence = [self.a]
        fraction = self.d

        while fraction is not None:
            sequence.append(fraction.a)
            fraction = fraction.d
        return sequence
    
    def expand(self, times:int = 1) -> 'ContinuedFraction':
        """
        Expand a continued fraction with the next term in the sequence

        Arguments:
        times -- The amount of times to expand
        """

        epsilon = 10**-30
        recurse = self
        while recurse.d is not None:
            recurse.inf = False # Can't have an infinite fraction in the middle
            recurse = recurse.d
            
        for t in range(times):
            if recurse.r < epsilon:
                return
            
            recurse.append(ContinuedFraction(Decimal('1.0')/Decimal(str(recurse.r)), depth=1))
            
            recurse = recurse.d

    def append(self, fraction:'ContinuedFraction') -> 'ContinuedFraction':
        """
        Append a continued fraction to this fraction

        Arguments:
        fraction -- Continued fraction to be appended
        """

        recurse = self
        while recurse.d is not None:
            recurse.inf = False # Can't have an infinite fraction in the middle
            recurse = recurse.d

        recurse.d = fraction
        return self
    
class ContinuedFractionSquareRootCanonicalForm:
    def __init__(self, N):
        self.N = N
        self.m = 0
        self.d = 1
        self.a = floor(sqrt(N))
        
    # https://en.wikipedia.org/wiki/Periodic_continued_fraction#Canonical_form_and_repetend
    def iterate(self):
        self.m = self.d*self.a - self.m
        self.d = floor((self.N - self.m**2)/self.d)
        self.a = floor((sqrt(self.N) + self.m)/self.d)
        return self
        
    def __eq__(self,other):
        return self.N == other.N and self.m == other.m and self.d == other.d and self.a == other.a
        
    def copy(self):
        other = ContinuedFractionSquareRootCanonicalForm(self.N)
        other.m = self.m
        other.d = self.d
        other.a = self.a
        return other