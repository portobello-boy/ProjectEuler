import sys
from MathLib import fraction

BOUND = 1000

def main():
    # Create infinite continued fraction for sqrt(2)
    cf = fraction.ContinuedFraction(addend=1, numerator=1, denominator=fraction.ContinuedFraction(addend=2, numerator=1, infinite=True))

    count = 0
    for d in range(BOUND):
        f = cf.evaluateFraction(depth=d)
        if len(str(f.n)) > len(str(f.d)):
            # print(d+1, f.str())
            count += 1
    print(count)

sys.setrecursionlimit(1500)
if __name__ == "__main__":
    main()