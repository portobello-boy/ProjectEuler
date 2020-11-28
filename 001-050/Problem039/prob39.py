import math

bound = 1001
perimeterCounts = dict()

def genPthagTriples(perimeterMap, bound):
    for m in range(2, bound):
        for n in range(m-1, 0, -1):
            if not math.gcd(m, n) == 1:
                continue
            
            p = 2*m*(m+n)
            k = 1

            while k*p < bound:
                perimeterMap[k*p] = perimeterMap.get(k*p, 0) + 1
                k += 1

genPthagTriples(perimeterCounts, bound)
print(perimeterCounts)
print(max(perimeterCounts, key=perimeterCounts.get))