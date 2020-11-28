bound = 9999 # 2*bound will have 5 digits, so concatenated with 1*bound will have 9 digits

def isPandigital(number):
    digits = [1, 2, 3, 4, 5, 6, 7, 8, 9]
    return sorted(number) == digits

def listToNum(l):
    return sum([val * 10**(ind) for ind, val in enumerate(l[::-1])])

def getLargestPandigitalMultiple():
    largestNum = 0
    largestProd = 0
    largestN = 1
    for num in range(bound):
        n = 1
        digitList = []
        while len(digitList) < 9:
            digitList.extend([int(i) for i in str(n * num)])
            n += 1
        if not isPandigital(digitList):
            continue

        val = listToNum(digitList)
        if val > largestProd:
            largestNum = num
            largestProd = val
            largestN = n-1
    return largestProd, largestNum, largestN

print(getLargestPandigitalMultiple())