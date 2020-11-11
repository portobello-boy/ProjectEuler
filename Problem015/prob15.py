import sys
from math import comb

height = int(sys.argv[1])
width = int(sys.argv[2])
print(comb(height+width, height))