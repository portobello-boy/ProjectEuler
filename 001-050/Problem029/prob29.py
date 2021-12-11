# def genSet(boundA, boundB):
#     seq = set()
#     for a in range(2, boundA+1):
#         for b in range(2, boundB+1):
#             seq.add(a**b)

#     return seq

# s = genSet(100, 100)
# print(len(s))

from itertools import product

INCLUSIVEBOUND = 100

perms = list(product(range(2, INCLUSIVEBOUND+1), range(2, INCLUSIVEBOUND+1)))
powers = set([p[0] ** p[1] for p in perms])

print(len(powers))
