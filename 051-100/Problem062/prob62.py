"""
The cube, 41063625 (3453), can be permuted to produce two other cubes: 56623104 (3843) and 66430125 (4053). In fact, 41063625 is the smallest cube which has exactly three permutations of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits are cube.
"""

NUM_PERMUTATIONS = 6
initialN = 345

cubes = {}
n = initialN
while True:
    key = ''.join(sorted(str(n**3)))
    cubes[key] = cubes[key] + [n**3] if key in cubes else [n**3]

    if len(cubes[key]) == NUM_PERMUTATIONS:
        print(min(cubes[key]))
        print(cubes[key])
        exit()

    n += 1
