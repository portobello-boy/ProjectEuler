from itertools import product
import sympy.ntheory as nt

primeFamilySize = 8

# Borrowed from https://www.geeksforgeeks.org/python-character-replacement-combination/
def genFormats(num):
    replacementDict = dict()
    for d in range(0, 10):
        replacementDict[str(d)] = [str(d), '*']

    formats = set()
    for sub in [zip(replacementDict.keys(), c) for c in product(*replacementDict.values())]:
        f = str(num)
        for rep in sub:
            f = f.replace(*rep)
        formats.add(f)
    return list(formats)

def genSubstitutions(string):
    nums = []
    if string[-1] == '*': # Last digit cannot be 0, 2, 4, 5, or 6
        cpy = string
        nums = [int(cpy.replace('*', str(d))) for d in [1, 3, 7, 9]]
    elif string[0] == '*':
        cpy = string
        nums = [int(cpy.replace('*', str(d))) for d in list(range(1, 10))]
    else:
        cpy = string
        nums = [int(cpy.replace('*', str(d))) for d in list(range(0, 10))]
    return nums

def haveSameDigits(numList):
    for i in range(len(numList)-1):
        if not len(nt.digits(numList[i])[1:]) == len(nt.digits(numList[i+1])[1:]):
            return False
    return True

def findPrimeFamily(size):
    visitedPrimes = set()
    primeFamily = set()
    p = 2

    while True:
        if p in visitedPrimes:
            p = nt.nextprime(p, 1)
            continue
        formats = genFormats(p) # List of strings
        for f in formats:
            primeFamily.clear()
            subs = genSubstitutions(f) # List of ints

            for s in subs:
                if nt.isprime(s):
                    visitedPrimes.add(s)
                    primeFamily.add(s)
                    
            # if len(primeFamily) == size and haveSameDigits(list(primeFamily)):
            if len(primeFamily) == size:
                return (p, f, sorted(list(primeFamily)))
        p = nt.nextprime(p, 1)

print(findPrimeFamily(primeFamilySize))