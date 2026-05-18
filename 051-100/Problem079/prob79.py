# 1. Find left-most digit (won't be in position 1 or 2 for any entry)
# 2. Find right-most digit (won't be in position 0 or 1)
# 3. Permute remaining digits between them until all constraints are satisfied
from collections import defaultdict
from itertools import permutations

constraints = defaultdict(list)
counts = defaultdict(dict)
valid_digits = set()

sum = 0
passkey_length = 0
with open("./051-100/Problem079/keylog.txt") as readFile:
    for n in readFile: 
        passkey_length = max(passkey_length, len(n) - 1)
        for i, v in enumerate(n):
            if v == '\n':
                continue
            if i not in counts:
                counts[i] = defaultdict(int)
            
            valid_digits.add(int(v))
            counts[i][int(v)] += 1
        constraints[int(n[0])].append(n[1:3])

print(passkey_length)
print(constraints)
print(counts)

first_digit = -1
last_digit = -1

for d1 in counts[0]:
    valid = True
    for idx_count in counts:
        if idx_count == 0:
            continue
        # print(f'checking {d1} in {counts[idx_count]}')
        if d1 in counts[idx_count]:
            # print(f'{d1} found in {counts[idx_count]}')
            valid = False
            continue
    if valid:
        first_digit = d1
        break
    
for d1 in counts[passkey_length-1]:
    valid = True
    for idx_count in counts:
        if idx_count == passkey_length-1:
            continue
        # print(f'checking {d1} in {counts[idx_count]}')
        if d1 in counts[idx_count]:
            # print(f'{d1} found in {counts[idx_count]}')
            valid = False
            continue
    if valid:
        last_digit = d1
        break

print(first_digit)
print(last_digit)

valid_digits.remove(first_digit)
valid_digits.remove(last_digit)
print(valid_digits)

def validate_constraints(password: str, constraints: dict) -> bool:
    for d in password:
        if int(d) not in constraints:
            continue
        for constraint in constraints[int(d)]:
            indices = []
            for c in constraint:
                indices.append(password.find(c))
            if indices != sorted(indices):
                return False
    return True

print(validate_constraints('531278', {3: ['17'], 5: ['82']}))

valid_passwords = set()

for p in permutations(list(valid_digits)):
    str_p = [str(d) for d in p]
    password = str(first_digit) + ''.join(str_p) + str(last_digit)
    # print(password)
    if validate_constraints(password, constraints):
        valid_passwords.add(int(password))
        
print(min(valid_passwords))