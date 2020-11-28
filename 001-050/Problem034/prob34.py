from math import factorial

# When searching for an upper bound, keep in mind that if 
# n is a natural number with d digits, then 
# 10^(d-1) <= n <= d*9! fails to hold with d >= 8. So the
# upper bound for our problem is 7 * 9!
def getFactorialNums():
    upperBound = 7 * factorial(9)
    factorialNums = set()
    for num in range(3, upperBound):
        digSum = sum(factorial(int(d)) for d in str(num))
        if digSum == num:
            factorialNums.add(num)
    return factorialNums

l = getFactorialNums()
print(l)
print(sum(l))