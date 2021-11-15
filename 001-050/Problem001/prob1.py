"""

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
"""

from MathLib.numberTheory import sumOfMultsOfN

INCLUSIVE_BOUND = 999

def main():
    print(sumOfMultsOfN(3, INCLUSIVE_BOUND) + sumOfMultsOfN(5, INCLUSIVE_BOUND) - sumOfMultsOfN(3 * 5, INCLUSIVE_BOUND))

if __name__ == "__main__":
    main()