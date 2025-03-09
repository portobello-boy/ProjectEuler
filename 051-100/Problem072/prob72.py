from MathLib.numberTheory import primeSieve

# BOUND = 8
BOUND=10**6

# Totient Sieve
totients = [i for i in range(BOUND+1)]

for p in primeSieve(BOUND):
    totients[p] = p-1

    k = 2
    while k * p <= BOUND:
        totients[k*p] -= totients[k*p]//p
        k += 1

print(sum(totients[2:]))