import sympy.ntheory as nt

primeList = [2]

def isGoldbachExample(n):
    global primeList
    # Expand primeList until it includes all primes p < n
    while primeList[-1] < n:
        p = primeList[-1]
        primeList.append(nt.nextprime(p, 1))

    for p in primeList:
        if p > n:
            break
        q = (n-p)/2
        if q**0.5 == int(q**0.5):
            return True
    return False

def findGoldbachCE(start):
    n = start
    while True:
        if not nt.isprime(n):
            if not isGoldbachExample(n):
                return n
        n += 2

print(findGoldbachCE(33))