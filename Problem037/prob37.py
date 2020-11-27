import sympy.ntheory as nt

def isTruncatablePrime(num, dir = 1):
    numArr = [int(i) for i in str(num)[::dir]]
    for n in range(len(numArr)):
        val = sum([val * 10**(ind) for ind, val in enumerate(numArr[n:][::-dir])])
        if not nt.isprime(val):
            return False
    return True

def isBiTruncatablePrime(num):
    return isTruncatablePrime(num) and isTruncatablePrime(num, -1)

def getTruncatablePrimes():
    truncatablePrimes = set()
    num = 9 # No single digit primes are considered truncatable primes
    while len(truncatablePrimes) < 11:
        num += 2
        if not nt.isprime(num):
            continue
        if isBiTruncatablePrime(num):
            truncatablePrimes.add(num)

    return truncatablePrimes        

tp = getTruncatablePrimes()
print(tp)
print(sum(tp))