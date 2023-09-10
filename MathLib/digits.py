from typing import List
import collections

def numDigits(n:int) -> int:
    return len(str(n))

# Return true if the number is palindromic
def isPalindrome(n:int) -> bool:
    return str(n) == str(n)[::-1]

def isPermutation(n:int, m:int) -> bool:
    d = collections.defaultdict(int)
    for x in str(n):
        d[x] += 1
    for x in str(m):
        d[x] -= 1
    return not any(d.values())

def getReversedNumber(n:int) -> int:
    return int(str(n)[::-1])

def digitProduct(n:int) -> int: # type: ignore
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

def concatDigits(n:int, m:int) -> int:
    """
    Given integers n and m, concatenage the digits of m to n
    """
    return int(''.join([str(n), str(m)]))