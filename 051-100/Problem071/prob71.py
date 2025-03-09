from math import gcd

BOUND = 10**6
TARGET = 3/7
n, d, gn, gd = 0, 1, 0, 1

while d < BOUND:
    while n/d < 3/7:
        if gcd(n, d) == 1 and n/d > gn/gd:
            gn, gd = n, d
        n += 1
    d += 1

print(TARGET)
print(gn)