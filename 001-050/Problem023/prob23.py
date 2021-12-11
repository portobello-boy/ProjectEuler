"""
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
"""

from math import sqrt
from MathLib.numberTheory import isAbundant, properDivisorList

BOUND = 28123

# for i in range(1, bound):
#     sumOfAbundants = False
#     for abundant in abundantNums:
#         if i-abundant in abundantNums:
#             sumOfAbundants = True
#             break
#     if not sumOfAbundants:
#         notAbundantSum += i

# print(notAbundantSum)

def main():
    abundantNums = set([n for n in range(12, BOUND) if isAbundant(n)])
    nonAbundantSum = 0

    for n in range(1, BOUND):
        sumOfAbundants = False
        for a in abundantNums:
            if a >= n:
                break
            if n-a in abundantNums:
                sumOfAbundants = True
                break
        if not sumOfAbundants:
            nonAbundantSum += n
    
    print(nonAbundantSum)


if __name__ == "__main__":
    main()
