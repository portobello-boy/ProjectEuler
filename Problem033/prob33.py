import math

# This should 'cancel' digits for any digit-length fraction
def cancel(num, den):
    numArr = [int(i) for i in str(num)]
    denArr = [int(i) for i in str(den)]

    commonDigits = []

    for num in numArr:
        if num in denArr:
            commonDigits.append(num)

    for num in commonDigits:
        if num in numArr:
            numArr.remove(num)
        if num in denArr:
            denArr.remove(num)

    newNum = sum([val * 10**(ind) for ind, val in enumerate(numArr[::-1])])
    newDen = sum([val * 10**(ind) for ind, val in enumerate(denArr[::-1])])

    return (newNum, newDen)

def isNonTrivCancellation(num, den):
    if num >= den:
        return False    # Fraction must be less than 1
    elif num % 10 == 0 and den % 10 == 0:
        return False    # Both are divisible by 10, this is trivial

    cancelledFrac = cancel(num, den)
    if cancelledFrac[1] == 0:
        return False    # Can't divide by 0

    if cancelledFrac[0] / cancelledFrac[1] == num / den:
        return True

    return False

def getNonTrivFractions():
    nonTrivFractions = set()

    # This generates pairs of 2-digit fractions with a common digit
    for i in range(1, 10):
        for j in range(0, 10):
            numerator = i*10 + j

            for k in range(0, 10):
                denominator = i*10 + k
                if isNonTrivCancellation(numerator, denominator):
                    nonTrivFractions.add((numerator, denominator))

                denominator = k*10 + i
                if isNonTrivCancellation(numerator, denominator):
                    nonTrivFractions.add((numerator, denominator))

                denominator = j*10 + k
                if isNonTrivCancellation(numerator, denominator):
                    nonTrivFractions.add((numerator, denominator))

                denominator = k*10 + j
                if isNonTrivCancellation(numerator, denominator):
                    nonTrivFractions.add((numerator, denominator))

    return nonTrivFractions

def fractionProduct(fractions):
    prod = 1
    for frac in fractions:
        prod *= frac[0]/frac[1]
    return prod

print(fractionProduct(getNonTrivFractions()))