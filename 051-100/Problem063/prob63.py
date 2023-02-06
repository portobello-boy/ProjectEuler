"""
The 5-digit number, 16807=75, is also a fifth power. Similarly, the 9-digit number, 134217728=89, is a ninth power.

How many n-digit positive integers exist which are also an nth power?
"""

base = 1
n = 1
count = 0

while base < 10: # Once base = 10, any n doesn't work
    while n <= len(str(base**n)):
        if len(str(base**n)) == n:
            count += 1
            print(base, n, base**n, count)
        n += 1
    n = 1
    base += 1

print(count)