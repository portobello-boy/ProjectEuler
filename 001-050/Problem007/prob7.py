"""
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
"""

from MathLib.numberTheory import nextPrime
from MathLib.constants import firstPrimes

COUNT = 10001

def main():
    curPrime = firstPrimes[-1]
    ind = len(firstPrimes)
    while ind != COUNT:
        curPrime = nextPrime(curPrime)
        ind += 1
    print(curPrime)

if __name__ == "__main__":
    main()