"""
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

    1634 = 1^4 + 6^4 + 3^4 + 4^4
    8208 = 8^4 + 2^4 + 0^4 + 8^4
    9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
"""

power = 5

def sumOfPowers(num, pow):
    return sum([int(c)**power for c in str(num)])

def getNarcissisticNums(pow):
    upperBound = pow * (9**pow)
    narcissistNums = set()
    for num in range(2, upperBound):
        digSum = sumOfPowers(num, pow)
        if digSum == num:
            narcissistNums.add(num)
    return narcissistNums

l = getNarcissisticNums(power)
print(l)
print(sum(l))