whichDigits = [1, 10, 100, 1000, 10000, 100000, 1000000]

def genChampernowneConstantString(length):
    constant = ""
    n = 1
    while len(constant) < length:
        constant += str(n)
        n += 1
    return constant

def getChampDigitProd(digitList):
    c = genChampernowneConstantString(max(whichDigits))
    prod = 1
    for d in digitList:
        prod *= int(c[d-1])
    return prod

print(getChampDigitProd(whichDigits))