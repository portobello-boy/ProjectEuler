from MathLib.numberTheory import multiplesInRange

INCLUSIVE_BOUND = 999

def main():
    print(multiplesInRange(3, INCLUSIVE_BOUND) + multiplesInRange(5, INCLUSIVE_BOUND) - multiplesInRange(3 * 5, INCLUSIVE_BOUND))

if __name__ == "__main__":
    main()