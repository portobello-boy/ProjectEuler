import sympy.ntheory as nt

bound = 1000000
primes = list(nt.sieve.primerange(2, bound))

# def getLongestSequenceForPrimeIndex(primeIndex):
#     """ This is super slow since it examines windows for every prime """
#     global primes
#     p = primes[primeIndex]
#     primeSubset = primes[:primeIndex+1]
#     # longestSeq = []

#     windowL = 0
#     windowR = 1
#     windowSum = sum(primeSubset[windowL:windowR])

#     # print(p, primeSubset)
#     while windowSum != p:
#         # print(windowL, windowR, primeSubset[windowL:windowR], windowSum, p)
#         if windowSum < p and windowR < len(primeSubset):
#             windowR += 1
#         elif windowSum > p and windowL < len(primeSubset):
#             windowL += 1
#         windowSum = sum(primeSubset[windowL:windowR])

#     return primeSubset[windowL:windowR]

def getLongestSequence():
    global bound
    global primes
    longestSeq = []
    prime = 0
    # Iterate over all primes which could have a longest sequence
    for i in range(len(primes) - len(longestSeq)):
        # Iterate over primes after i which could be part of a sequence
        for j in range(i+len(longestSeq)+1, len(primes)):
            primeSubset = primes[i:j+1]
            subsetSum = sum(primeSubset)
            if subsetSum > bound or subsetSum < prime:
                break
            if nt.isprime(subsetSum) and len(primeSubset) > len(longestSeq):
                longestSeq = primeSubset
                prime = subsetSum

    # for i in range(len(primes)-1, 0, -1):
        # if len(longestSeq) > len(primes[:i]):
        #     continue
        # seq = getLongestSequenceForPrimeIndex(i)
        # if len(seq) > len(longestSeq):
        #     longestSeq = seq
        #     prime = primes[i]
        
    return longestSeq, prime

# print(primes)
l = getLongestSequence()
print(l[0], l[1], len(l[0]))