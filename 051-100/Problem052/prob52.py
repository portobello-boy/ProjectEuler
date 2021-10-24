def sortedDigits(num):
    return sorted([d for d in str(num)])

def isPermuted(num):
    digits = sortedDigits(num)
    for n in range(2, 7):
        if sortedDigits(n*num) != digits:
            return False
    return True

num = 2
while not isPermuted(num):
    num += 1

print(num)