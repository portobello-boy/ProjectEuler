"""
Consider quadratic Diophantine equations of the form:

x2 – Dy2 = 1

For example, when D=13, the minimal solution in x is 6492 – 13×1802 = 1.

It can be assumed that there are no solutions in positive integers when D is square.

By finding minimal solutions in x for D = {2, 3, 5, 6, 7}, we obtain the following:

32 – 2×22 = 1
22 – 3×12 = 1
92 – 5×42 = 1
52 – 6×22 = 1
82 – 7×32 = 1

Hence, by considering minimal solutions in x for D ≤ 7, the largest x is obtained when D=5.

Find the value of D ≤ 1000 in minimal solutions of x for which the largest value of x is obtained.
"""

from MathLib.fraction import ContinuedFraction
from math import sqrt
from decimal import Decimal

def main():
    UPPER_BOUND = 1000
    maxX = 0
    maxD = 0

    perfectSquares = set([i**2 for i in range(int(sqrt(UPPER_BOUND))+1)])
    
    for D in range(2, UPPER_BOUND+1):
        if D in perfectSquares:
            continue
        
        # Todo - clean up CF constructor arguments to not need Decimal here
        cf = ContinuedFraction(Decimal(D).sqrt(), depth=1)
        ptr = cf
        f = cf.evaluateFraction(depth=-1)
        
        while (f.n**2 - (D * f.d**2)) != 1:
            ptr.expand()
            ptr = ptr.d
            f = cf.evaluateFraction(depth=-1)
        
        if f.n > maxX:
            maxX = f.n
            maxD = D
        
    print(maxX)
    print(maxD)


if __name__ == "__main__":
    main()
    # test()