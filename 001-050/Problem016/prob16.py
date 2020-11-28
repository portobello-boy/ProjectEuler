import math
import sys

power = int(sys.argv[1])
num = 2 ** power

print(sum([int(d) for d in str(num)]))