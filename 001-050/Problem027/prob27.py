import math
import sympy.ntheory as nt

def genFunction(a, b):
    def f(n):
        return n**2 + a*n + b
    return f

def countConsecutivePrimes(a, b):
    f = genFunction(a, b)
    n = 0

    while nt.isprime(f(n)):
        n += 1

    return n

def bestCoefficientProduct(boundA, boundB):
    maxProd = 0
    maxConsecutivePrimes = 0

    for a in range(boundA):
        for b in range(boundB+1):
            if not nt.isprime(b):
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

print(bestCoefficientProduct(1000, 1000))