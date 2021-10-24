from MathLib.numberTheory import fibonacci
from MathLib.constants import golden

BOUND_INCLUSIVE = 4 * 10**6

def main():
    sum = 0
    index = 3

    cur = fibonacci(index) # F_3 = 2
    while cur < BOUND_INCLUSIVE:
        sum += cur
        index += 3
        cur = fibonacci(index)
    
    print(sum)

if __name__ == "__main__":
    main()