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