from MathLib.fraction import ContinuedFraction

TERM = 100

sequence = [2, 1]
k = 1
while len(sequence) < TERM:
    sequence += [2 * k, 1, 1]
    k += 1
    
sequence = sequence[:TERM]

print(sequence)
print(len(sequence))

frac = ContinuedFraction.generateFromSequence(sequence)
evalFrac = frac.evaluateFraction(depth=TERM)

print(evalFrac.str())
print(sum([int(k) for k in str(evalFrac.n)]))