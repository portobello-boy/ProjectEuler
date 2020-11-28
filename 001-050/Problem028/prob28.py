w = 1001

def diagonalSpiralSum(width):
    total = 1
    step = 2
    num = 1

    while num < width**2:
        for _ in range(4):
            num += step
            total += num
        step += 2

    return total

print(diagonalSpiralSum(w))