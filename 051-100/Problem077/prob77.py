from MathLib.constants import firstPrimes
from MathLib.numberTheory import nextPrime
from cachetools import cached
from cachetools.keys import hashkey

primes = sorted(firstPrimes)

num = 11

# @cached(
#     cache={},
#     key=lambda num, values, partitions: hashkey(num)
# )
def partition(num: int, values: list[int], partitions: list[int]) -> int:
    if num < 0:
        return 0
    if num == 0:
        print('partition:', partitions)
        return 1
    if num == 1:
        return 0
    if len(values) == 0:
        return 0
    
    return partition(num, values[:-1], partitions) + partition(num - values[-1], values, partitions + [values[-1]])

count = 0
while count < 5000:
    num += 1
    if num > primes[-1]:
        primes += [nextPrime(primes[-1])]
    count = partition(num, primes, [])
    
print(num, count)