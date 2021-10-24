from sympy import primefactors

def findIntsWithFactors(numInts, numFactors):
    num = 1
    while True:
        ints = list([num + i for i in range(numInts)])

        count = 0
        for val in ints:
            valFactors = primefactors(val)
            if len(valFactors) == numFactors:
                count += 1
        if count == numInts:
            return ints

        num += 1

print(findIntsWithFactors(3, 4))