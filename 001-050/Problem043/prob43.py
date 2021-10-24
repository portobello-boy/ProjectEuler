from itertools import permutations
import sympy.ntheory as nt

def listToNum(l):
    return sum([val * 10**(ind) for ind, val in enumerate(l[::-1])])

def getPandigitalList():
    digits = [*range(0, 10)]
    pandigitalPermutations = list(permutations(digits, len(digits)))
    return pandigitalPermutations

def hasSubstringDivisibility(numList):
    for i in range(1, len(numList)-2):
        p = nt.prime(i)
        num = listToNum(numList[i:i+3])
        if not num % p == 0:
            return False
    return True

def sumOfPan():
    l = getPandigitalList()
    sum = 0
    for lst in l:
        if hasSubstringDivisibility(lst):
            sum += listToNum(lst)
    return sum

print(sumOfPan())