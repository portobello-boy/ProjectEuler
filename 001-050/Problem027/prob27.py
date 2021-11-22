"""
See https://projecteuler.net/problem=27
"""

from MathLib.polynomials import generatePolynomial
from MathLib.numberTheory import isPrime

def countConsecutivePrimes(a, b):
    f = generatePolynomial([1, a, b])

    x = 0
    while isPrime(f(x)):
        x += 1

    return x

def bestCoefficientProduct(boundA, boundB):
    maxProd = 0
    maxConsecutivePrimes = 0

    for a in range(boundA):
        for b in range(boundB+1):
            if not isPrime(b):
                continue
            if (x := countConsecutivePrimes(a, b)) > maxConsecutivePrimes:
                maxConsecutivePrimes = x
                maxProd = a*b
            if (x := countConsecutivePrimes(-a, b)) > maxConsecutivePrimes:
                maxConsecutivePrimes = x
                maxProd = -a*b
            if (x := countConsecutivePrimes(a, -b)) > maxConsecutivePrimes:
                maxConsecutivePrimes = x
                maxProd = a*-b
            if (x := countConsecutivePrimes(-a, -b)) > maxConsecutivePrimes:
                maxConsecutivePrimes = x
                maxProd = -a*-b

    return maxProd

def main():
    print(bestCoefficientProduct(1000, 1000))

if __name__ == "__main__":
    main()