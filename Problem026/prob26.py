import math
import sympy.ntheory as nt

bound = 1000

def GCD (a, b ) : 
    if (b == 0 ) : 
        return a 
    return GCD( b, a % b )

def multiplicativeOrder(A, N) : 
    if (GCD(A, N ) != 1) : 
        return -1

    result = 1
    K = 1
    while (K < N) : 
        result = (result * A) % N
  
        if (result == 1) : 
            return K

        K = K + 1
      
    return -1

d = 1
cycleLength = 1

for i in range(2, bound):
    if not nt.isprime(i):
        continue
    order = multiplicativeOrder(10, i)
    if order != -1 and order > cycleLength:
        cycleLength = order
        d = i

print(d, cycleLength)