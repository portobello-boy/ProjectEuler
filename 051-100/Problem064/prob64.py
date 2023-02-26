from math import sqrt
from MathLib.fraction import ContinuedFractionSquareRootCanonicalForm as Expansion

UPPER_BOUND = 10000

perfectSquares = set([i**2 for i in range(int(sqrt(UPPER_BOUND))+1)])

oddCount = 0

for num in range(UPPER_BOUND + 1):
    if num in perfectSquares:
        continue
    
    e = Expansion(num)
    
    periodCycle = []
    
    while e.iterate() not in periodCycle:
        periodCycle.append(e.copy())
            
    if len(periodCycle) % 2 == 1:
        oddCount += 1

print(oddCount)