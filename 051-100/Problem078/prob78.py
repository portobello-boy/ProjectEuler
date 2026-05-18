from MathLib.numberTheory import partitionCount

i = 1
while (p := partitionCount(i)) % 10**6 != 0:
    i += 1
    if i % 10 == 0:
        print('i:', i, 'p:', p)
    
print(i, p)

# https://nicf.net/articles/generating-functions-partitions/
# https://www.whitman.edu/mathematics/cgt_online/book/section03.03.html
# https://pages.uoregon.edu/koch/PentagonalNumbers.pdf
