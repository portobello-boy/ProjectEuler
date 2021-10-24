import csv

def getScore(name):
    return sum([ord(c)-64 for c in name])

def getTotalScore(names):
    return sum((ind+1) * getScore(name) for ind, name in enumerate(names))

with open('names.txt') as csvfile:
    reader = csv.reader(csvfile,delimiter=',',quotechar='"')
    names = reader.__next__()
    print(getTotalScore(sorted(names)))