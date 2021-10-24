def makeChange(coins,amount):
    if amount < 0:
        return 0
    if len(coins) == 1 or amount == 0:
        return 1
    return makeChange(coins[:-1], amount) + makeChange(coins, amount - coins[-1])

print(makeChange([1,2,5,10,20,50,100,200], 200))