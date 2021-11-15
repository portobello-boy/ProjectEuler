"""
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
"""

from MathLib.numberTheory import divisorList, isPrime

NUM = 600851475143

def main():
    divisors = divisorList(NUM)

    for i in reversed(divisors):
        if isPrime(i):
            print(i)
            return

if __name__ == "__main__":
    main()