from itertools import permutations

NODE_COUNT = 10 # 6 for test
DIGITS = set([i for i in range(1, NODE_COUNT+1)])

SOLUTIONS = set()
# SOLUTION_FORMAT = [0, 1, 2, 3, 2, 4, 5, 4, 1]
SOLUTION_FORMAT = [0, 1, 2, 3, 2, 6, 4, 6, 7, 5, 7, 8, 9, 8, 1]

# initialValues = permutations(DIGITS, 4)
initialValues = permutations(DIGITS, 6)

def getSolutionString(vals):
    # outerNodes = [vals[0], vals[3], vals[5]]
    outerNodes = [vals[0], vals[3], vals[4], vals[5], vals[9]]
    offset = outerNodes.index(min(outerNodes)) * 3
    rotatedFormat = SOLUTION_FORMAT[offset:] + SOLUTION_FORMAT[:offset]
    return ''.join([str(vals[i]) for i in rotatedFormat])

for vals in list(initialValues):
    v = sum(vals[:3])
    
    # e = v - vals[3] - vals[2]
    # if e not in DIGITS or e in vals:
    #     continue
    
    # vals += (e, )
    
    # f = v - vals[1] - e
    # if f not in DIGITS or f in vals:
    #     continue
    
    # vals += (f, )

    g = v - vals[3] - vals[2]
    if g not in DIGITS or g in vals:
        continue
    
    vals += (g, )
    
    h = v - vals[4] - g
    if h not in DIGITS or h in vals:
        continue
    
    vals += (h, )
    
    i = v - vals[5] - h
    if i not in DIGITS or i in vals:
        continue
    
    vals += (i, )
    
    j = v - vals[1] - i
    if j not in DIGITS or j in vals:
        continue
    
    vals += (j, )
    SOLUTIONS.add(getSolutionString(vals))
    
print(SOLUTIONS)
print(max(SOLUTIONS))