def getDigitsOfSelfPower(digits, power):
    s = str(sum([i**i for i in range(1, power+1)]))
    return s[len(s)-digits:]

print(getDigitsOfSelfPower(10, 1000))