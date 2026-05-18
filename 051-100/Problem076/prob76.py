TARGET = 100

def partitions(num, I=1):
    yield (num,)
    for i in range(I, num//2 + 1):
        for p in partitions(num-i, i):
            yield (i,) + p
    
print(sum(1 for p in partitions(TARGET) if len(p) != 1))