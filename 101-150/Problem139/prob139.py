# Written by Cameron Haddock & Daniel Millson
# Submitted as a solution to Project Euler's Problem 139


# Given that the perimeter of the right triangle is less than one-hundred million, how many Pythagorean triangles would allow such a tiling to take place?

from math import gcd

MAX_PERIMETER = 10**8

UPPER_M = 5774
LOWER_M = 1
UPPER_N = 4083
LOWER_N = 1

valid_triangles = 0

for n in range(LOWER_N, UPPER_N):
    for m in range(n + 1, UPPER_M, 2):
        if gcd(m, n) != 1:
            continue

        a = m**2 - n**2
        b = 2*m*n
        c = m**2 + n**2

        if c % abs(a-b) != 0:
            continue

        sum_ = a + b + c
        k = 1
        while k*sum_ < MAX_PERIMETER:
            ka, kb, kc = k*a, k*b, k*c
            
            difference = abs(ka - kb)

            if kc % difference == 0:
                valid_triangles += 1

            k += 1

print(valid_triangles)