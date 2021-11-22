"""
Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.

    [see numsDec.txt]
"""

def main():
    sum = 0
    with open("./numsDec.txt") as readFile:
        for n in readFile:
            sum += int(n)
    print(str(sum)[0:10])

if __name__ == "__main__":
    main()