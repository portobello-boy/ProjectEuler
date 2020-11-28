knownTriangleNums = set()
knownPentagonalNums = set()
knownHexagonalNums = set()

def t(n):
    return (1/2)*n*(n+1)

def p(n):
    return (1/2)*n*(3*n-1)

def h(n):
    return n*(2*n-1)

def isTriangleNum(num):
    global knownTriangleNums
    # Generate triangle nums until num is hit or surpassed
    n = len(knownTriangleNums)
    while True:
        n += 1
        knownTriangleNums.add(t(n))
        if t(n) > num:
            break

    if num in knownTriangleNums: # If num is a known triangle number
        return True
    return False

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

def isHexagonalNum(num):
    global knownHexagonalNums
    # Generate hexagonal nums until num is hit or surpassed
    n = len(knownHexagonalNums)
    while True:
        n += 1
        knownHexagonalNums.add(h(n))
        if h(n) > num:
            break

    if num in knownHexagonalNums: # If num is a known hexagonal number
        return True
    return False

def findCommonFigurativeNum(which):
    n = 1
    count = 0
    while True:
        hn = h(n)
        if isTriangleNum(hn) and isPentagonalNum(hn) and isHexagonalNum(hn):
            count += 1
            if count == which:
                return hn
        n += 1

print(findCommonFigurativeNum(4))