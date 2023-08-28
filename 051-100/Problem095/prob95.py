from MathLib.numberTheory import properDivisorList

BOUND = 1_000
seen: set[int] = set()
longestChain: set[int] = set()

for num in range(BOUND):
    loop: set[int] = set()
    next_ = num
    while True:
        if next_ in seen or next_ >= BOUND:
            break
        if next_ == num:
            print(num)
            if len(loop) > len(longestChain):
                longestChain = loop
            seen = seen | loop
            break
        if next_ in loop:
            break    
        next_ = sum(properDivisorList(num))
        loop.add(next_)

print(longestChain)
