import csv

knownTriangleNums = set()

def t(n):
    return (1/2)*n*(n+1)

def isTriangleNum(num):
    global knownTriangleNums
    # Generate triangle nums until num is hit or surpassed
    n = len(knownTriangleNums)+1
    while t(n) <= num:
        knownTriangleNums.add(t(n))
        n += 1

    if num in knownTriangleNums: # If num is a known triangle number
        return True
    return False

def wordValue(word):
    return sum([ord(c)-64 for c in word])

def getTriangleWordCount(words):
    count = 0
    for word in words:
        if isTriangleNum(wordValue(word)):
            count += 1
    return count

with open('words.txt') as csvfile:
    reader = csv.reader(csvfile,delimiter=',',quotechar='"')
    words = reader.__next__()
    print(getTriangleWordCount(sorted(words)))