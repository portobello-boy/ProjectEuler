from math import factorial as fact
from math import ceil

START = 1
MAX = 100
TARGET = 1000000

def combination(n, r):
    return fact(n)/(fact(r) * fact(n-r))

count = 0
for n in range(START, MAX+1):
    for r in range(0, n):
        if combination(n, r) > TARGET:
            count += 1

print(count)