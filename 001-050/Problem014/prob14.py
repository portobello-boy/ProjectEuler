"""
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
"""

from MathLib.numberTheory import evenOddFunctionSequence

BOUND = 10**6

def main():
    collatz = evenOddFunctionSequence(lambda n: n/2, lambda n: 3*n+1)
    
    longestN = 0
    longestNLength = 0

    for startN in range(1, BOUND):
        n = collatz(startN)
        seqLength = 0
        while n != 1:
            n = collatz(n)
            seqLength += 1
        if seqLength > longestNLength:
            longestN = startN
            longestNLength = seqLength
    
    print(longestN, longestNLength)

if __name__ == "__main__":
    main()