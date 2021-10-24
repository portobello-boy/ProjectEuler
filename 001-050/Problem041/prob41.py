from itertools import permutations
import sympy.ntheory as nt

def listToNum(l):
    return sum([val * 10**(ind) for ind, val in enumerate(l[::-1])])

def getLargestNPandigitalPrime(maxDigits):
    largest = 0
    digits = [*range(1, maxDigits+1)]
    pandigitalPermutations = list(permutations(digits, len(digits)))
    for panPerm in pandigitalPermutations:
        if panPerm[-1] % 2 == 0:
            continue
        num = listToNum(panPerm)
        if nt.isprime(num):
            largest = max(largest, num)
    return largest

def getLargestPandigitalPrime():
    largest = 0
    for d in range(1, 10):
        largest = max(largest, getLargestNPandigitalPrime(d))
    return largest

print(getLargestPandigitalPrime())