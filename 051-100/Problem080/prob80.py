from decimal import *

getcontext().prec = 105

BOUND = 100

print(sum([sum([int(d) for d in str(Decimal(n).sqrt())[:101] if (d != '.' and int(Decimal(n).sqrt()) != Decimal(n).sqrt())]) for n in range(1, BOUND+1)]))

# sum = 0
# for n in range(1, BOUND+1):
#     ns = Decimal(n).sqrt()
#     if ns == int(ns):
#         print(f'{ns} is an int')
#         continue
#     decimal_offset = str(ns).find('.') + 1
#     # print(str(ns)[:101])
#     for d in str(ns)[:101]:
#         if d == '.':
#             continue
#         sum += int(d)
# print(sum)