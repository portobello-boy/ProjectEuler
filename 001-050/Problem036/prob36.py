bound = 1000000

def isPalindrome(num):
    return num == num[::-1]

def isDoublePalindrome(num):
    return isPalindrome(str(num)) and isPalindrome("{0:b}".format(num))

def sumPalindrome(bound):
    total = 0
    for i in range(bound):
        if isDoublePalindrome(i):
            total += i
    return total

print(sumPalindrome(bound))