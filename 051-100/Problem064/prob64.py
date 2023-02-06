from math import sqrt, floor

UPPER_BOUND = 10000

perfectSquares = set([i**2 for i in range(int(sqrt(UPPER_BOUND))+1)])

class Expansion:
    def __init__(self, N):
        self.N = N
        self.m = 0
        self.d = 1
        self.a = floor(sqrt(N))
        
    # https://en.wikipedia.org/wiki/Periodic_continued_fraction#Canonical_form_and_repetend
    def iterate(self):
        self.m = self.d*self.a - self.m
        self.d = floor((self.N - self.m**2)/self.d)
        self.a = floor((sqrt(self.N) + self.m)/self.d)
        return self
        
    def __eq__(self,other):
        return self.N == other.N and self.m == other.m and self.d == other.d and self.a == other.a
        
    def copy(self):
        other = Expansion(self.N)
        other.m = self.m
        other.d = self.d
        other.a = self.a
        return other

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