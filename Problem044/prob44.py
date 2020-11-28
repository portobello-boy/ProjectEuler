knownPentagonalNums = set()

def p(n):
    return (1/2)*n*(3*n-1)

def isPentagonalNum(num):
    global knownPentagonalNums
    # Generate pentagonal nums until num is hit or surpassed
    n = len(knownPentagonalNums)
    while True:
        n += 1
        knownPentagonalNums.add(p(n))
        if p(n) > num:
            break

    if num in knownPentagonalNums: # If num is a known pentagonal number
        return True
    return False

def findMinimalPentagonalPair():
    global knownPentagonalNums

    j = 1
    k = 1
    while True:
        pj = p(j)
        pk = p(k)

        if isPentagonalNum(pj+pk) and isPentagonalNum(abs(pk-pj)):
            return abs(pk-pj)

        if j == k:
            j += 1
            k = 1
            continue
        k += 1

print(findMinimalPentagonalPair())