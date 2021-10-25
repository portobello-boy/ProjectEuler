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