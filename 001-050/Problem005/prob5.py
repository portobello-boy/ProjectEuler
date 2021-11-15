"""
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
"""

from MathLib.numberTheory import lcm

BOUND = 20

def lcmMulti(a, b):
    if b == 1:
        return lcm(a, b)
    return lcm(a, lcmMulti(b, b-1))

def main():
    print(lcmMulti(BOUND, BOUND-1))

if __name__ == "__main__":
    main()