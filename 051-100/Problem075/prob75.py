from MathLib.numberTheory import properDivisorList
import numpy as np
from math import sqrt, ceil
from collections import defaultdict

UPPER_BOUND = 1500000
# UPPER_BOUND = 120
sum_counts = defaultdict(int)

for m in range(2, 866):
    for n in range(m-1, 0, -2):
        if np.gcd(m, n) != 1:
            continue
        
        triple = (m**2 - n**2, 2*m*n, m**2 + n**2)
        
        k = 1
        while (triple_sum := k * sum(triple)) < UPPER_BOUND:
            sum_counts[triple_sum] += 1
            k += 1
            
total = 0
for k, v in sorted(sum_counts.items()):
    if v == 1:
        # print(k, ': ', v)
        total += 1
print(total)