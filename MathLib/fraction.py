from typing import Optional, Sequence
from MathLib.numberTheory import gcd

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
        self.inf = infinite

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

        if not depth == 0 and self.inf:
            # print("INFINITE")
            subFrac = Fraction(self.n) * self.evaluateFraction(depth = depth-1).reciprocal()
            return subFrac.addInt(self.a)
        
        if (depth == 0 and self.inf) or self.n == 0 or self.d is None:
            return Fraction(self.a)
            
        subFrac = Fraction(self.n) * self.d.evaluateFraction(depth = depth).reciprocal()
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