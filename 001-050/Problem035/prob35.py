import sympy.ntheory as nt

bound = 1000000

def getListRotations(l):
    rotations = list()
    for i in range(len(l)):
        rotations.append(l[i:] + l[:i])
    return rotations

def listToNum(l):
    return sum([val * 10**(ind) for ind, val in enumerate(l)])

def isCircularPrime(num):
    rotations = getListRotations([int(i) for i in str(num)])
    for rot in rotations:
        possiblePrime = listToNum(rot[::-1])
        if not nt.isprime(possiblePrime):
            return False
    return True

def getCircularPrimes():
    circularPrimes = {2, 3, 5, 7}
    for num in range(11, bound, 2):
        if isCircularPrime(num):
            circularPrimes.add(num)
    return circularPrimes

cp = getCircularPrimes()
print(cp)
print(len(cp))