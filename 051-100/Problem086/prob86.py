# Written by Cameron Haddock
# Submitted as a solution to Project Euler's Problem 85

from math import gcd, ceil

# A spider, S, sits in one corner of a cuboid room, measuring 6 by 5 by 3, and a fly, F, sits in the opposite corner. By travelling on the surfaces of the room the shortest "straight line" distance from S to F is 10 and the path is shown on the diagram.
# However, there are up to three "shortest" path candidates for any given cuboid and the shortest route doesn't always have integer length.
# It can be shown that there are exactly 2060 distinct cuboids, ignoring rotations, with integer dimensions, up to a maximum size of M by M by M, for which the shortest route has integer length when M = 100. This is the least value of M for which the number of solutions first exceeds two thousand; the number of solutions when M = 99 is 1975.
# Find the least value of M such that the number of solutions first exceeds one million.

count_paths = 0
target = 2000
M = 2
m = 2
n = 1
explored_triples = set() # (3,4,5), (5,12,13), (6,8,10), (7, 10, ?)...
valid_paths = set()

while count_paths < target:
    count_paths = 0
    # Increase the largest dimension of the cuboid
    M += 1

    # Generate all primitive pythagorean triples representing a path from opposite diagonals
    # Generate triple with hypotenuse <= maximum diagonal length of cuboid
    while n**2 + n**2 <= (M**2 + (2*M)**2)**0.5:
        a = m**2 - n**2
        b = 2*m*n
        c = m**2 + n**2

        print(a, b, c, "--", m, n)
        explored_triples.add((min(a, b), max(a, b), c))

        n += 2
        while gcd(m, n) != 1:
            n += 2

        if n > m:
            m += 1
            n = 1 + (m % 2)

    # print(explored_triples, n**2 + n**2, (M**2 + (2*M)**2)**0.5)
    # Scale known primitive triples so that c <= maximum diagonal length of cuboid
    new_triples = set()
    for triple in explored_triples:
        k = 2
        while k*triple[2] <= (M**2 + (2*M)**2)**0.5:
            new_triples.add(tuple(k*e for e in triple))
            k += 1
    # print(new_triples)
    explored_triples = explored_triples.union(new_triples)

    # Select from explored_triple the triples that fit in cuboid up to MxMxM
    for triple in explored_triples:
        if triple[0] <= M and triple[1] <= 2*M:
            valid_paths.add(triple)

    for path in valid_paths:
        count = 0
        for i in range(0, path[0]+1):
            if path[1]-i <= path[1] and i <= path[0]:
                count += 1
        print("Dimension:", M, "validPath:", path, "how many:", count)
        count_paths += count

# print(sorted(explored_triples))
# print(sorted(valid_paths))
print(count_paths)
print(M)

l = 2
test_count = 0

while test_count < target:
    l += 1
    for wh in range(3, 2*l + 1):
        if ((wh**2 + l**2)**0.5).is_integer():
            if (wh <= l):
                test_count += wh/2
            else:
                test_count += 1 + (l - (wh+1)/2)

print(test_count)
print(l)