def genSet(boundA, boundB):
    seq = set()
    for a in range(2, boundA+1):
        for b in range(2, boundB+1):
            seq.add(a**b)

    return seq

s = genSet(100, 100)
print(len(s))