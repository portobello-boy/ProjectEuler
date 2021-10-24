from num2words import num2words
import sys

bound = int(sys.argv[1])
count = 0
for i in range(1, bound+1):
    count += sum([1 for c in num2words(i) if c.isalnum()])
print(count)