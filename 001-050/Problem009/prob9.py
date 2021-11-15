"""
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
    a^2 + b^2 = c^2

For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
"""
from MathLib.numberTheory import gcd

TARGET = 1000

def generatePythagoreanTriple(targetSum):
    for m in range(2, targetSum):
        for n in range(m-1, 0, -2):
            if gcd(m, n) != 1:
                continue

            a = m**2 - n**2
            b = 2 * m * n
            c = m**2 + n**2

            sum = a + b + c

            if targetSum % sum == 0:
                k = targetSum / sum

                return int(a * b * c * k**3)


def main():
    print(generatePythagoreanTriple(TARGET))

if __name__ == "__main__":
    main()