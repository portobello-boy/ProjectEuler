from math import ceil
from MathLib.numberTheory import gcd

BOUND = 12000
LOWER_TARGET = 1/3
UPPER_TARGET = 1/2
d = 1
count = 0
fracs = set()

while d < BOUND+1:
    n = ceil(d * LOWER_TARGET)
    while n/d < UPPER_TARGET:
        if gcd(n, d) == 1 and n/d > LOWER_TARGET:
            count += 1
            fracs.add((n, d))
        n += 1
    d += 1

print(count)
# print(fracs)