from math import factorial

num = factorial(100)
string = str(num)
digSum = sum([int(d) for d in string])

print(digSum)