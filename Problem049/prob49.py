from itertools import permutations
import sympy.ntheory as nt

def listToNum(l):
    return sum([val * 10**(ind) for ind, val in enumerate(l[::-1])])

def getPrimePerms(numList):
    perms = list(permutations(numList, len(numList)))
    primes = set()
    for perm in perms:
        if perm[0] == 0:
            continue
        val = listToNum(perm)
        if nt.isprime(val):
            primes.add(val)
    return primes

def exhibitsProperty(termList):
    for i in range(0, len(termList)-2):
        if not abs(termList[i] - termList[i+1]) == abs(termList[i+1] - termList[i+2]):
            return False
    return True

def getDigitalPrimePerms(numDigits, numTerms):
    primeSequences = []
    encounteredPrimes = set()
    for n in range(10**(numDigits-1), 10**numDigits):
        if n in encounteredPrimes:
            continue
        if not nt.isprime(n):
            continue
        primePerms = getPrimePerms([int(i) for i in str(n)])
        encounteredPrimes.update(primePerms)
        
        sequenceList = list(permutations(primePerms, numTerms))
        for sequence in sequenceList:
            srtSq = sorted(sequence)
            if exhibitsProperty(srtSq) and not srtSq in primeSequences:
                primeSequences.append(srtSq)

    return primeSequences

print(getDigitalPrimePerms(4, 3))