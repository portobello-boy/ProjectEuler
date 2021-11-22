# Return true if the number is palindromic
from typing import List

def numDigits(n:int) -> int:
    return len(str(n))

def isPalindrome(n:int) -> bool:
    return str(n) == str(n)[::-1]

def digitProduct(n:int) -> int:
    return digitProduct([int(d) for d in str(n)])

def digitProduct(n:List[int]) -> int:
    prod = 1
    for d in n:
        prod *= d
    return prod

# Return the maximum product of s adjacent digits in n
def digitProductMax(n:int, s:int) -> int:
    numStr = [int(d) for d in str(n)]

    max = 0
    for i in range(0, len(numStr) - s + 1):
        prod = digitProduct(numStr[i:i+s])
        if prod > max:
            max = prod
    return max