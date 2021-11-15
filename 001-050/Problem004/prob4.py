"""
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
"""

from MathLib.digits import isPalindrome

def main():
    max = 0
    for m in range((10**3)-1, 100-1, -1):
        for n in range((10**3)-1, m-1, -1):
            if isPalindrome(m*n) and m*n > max:
                max = m*n
    print(max)

if __name__ == "__main__":
    main()