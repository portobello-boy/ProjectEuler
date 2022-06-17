"""
The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and concatenating them in any order the result will always be prime. For example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes, 792, represents the lowest sum for a set of four primes with this property.

Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.
"""

from itertools import combinations, permutations
from MathLib.digits import concatDigits
from MathLib.numberTheory import isPrime, nextPrime, primeSequence, primeSequenceBounded


NUM_PRIMES = 4
primes = primeSequence(NUM_PRIMES)

def checkConcatenation(primes):
    for perms in permutations(primes, 2):
        candidate = int(''.join(str(i) for i in perms))
        if not isPrime(candidate):
            return False
    return True

found = False
while not found:
    largestPrime = primes[-1]
    # print(largestPrime)

    candidatePrimes = set([largestPrime])
    for n in range(NUM_PRIMES):
        for p in primes[:len(primes)-1]:
            if p in candidatePrimes:
                continue

            candidatePrimes.add(p)
            if not checkConcatenation(candidatePrimes):
                candidatePrimes.remove(p)
                continue

            if len(candidatePrimes) == NUM_PRIMES:
                print(candidatePrimes)
                print(sum(candidatePrimes))
                found = True

    largestPrime = nextPrime(largestPrime)
    primes.append(largestPrime)