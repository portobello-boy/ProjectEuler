"""
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
"""
from MathLib.numberTheory import primeSequenceBoundedGenerator, primeSequenceGenerator

BOUND = 2 * 10**6

def main():
    print(sum(primeSequenceBoundedGenerator(BOUND)))

if __name__ == "__main__":
    main()