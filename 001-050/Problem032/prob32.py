def isPandigital(number):
    digits = ["1", "2", "3", "4", "5", "6", "7", "8", "9"]
    return sorted(list(number)) == digits

def getPandigitals():
    pandigitals = set()
    for i in range(1, 10):
        for j in range(1000, 10000):
            if isPandigital(str(i) + str(j) + str(i*j)):
                pandigitals.add(i*j)

    for i in range(10, 100):
        for j in range(100, 1000):
            if i*j >= 10000:
                break
            if isPandigital(str(i) + str(j) + str(i*j)):
                pandigitals.add(i*j)

    return pandigitals

print(sum(getPandigitals()))