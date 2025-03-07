LENGTH = 50
MIN_BLOCK_SIZE = 50

GAP_CACHE = {}

def solution_two(remainingSpaces):
    if remainingSpaces < MIN_BLOCK_SIZE:
        return 0
    
    if remainingSpaces in GAP_CACHE:
        return GAP_CACHE[remainingSpaces]
    
    total_combinations = 0
    for position in range(remainingSpaces):
        # iterate over block sizes valid for range
        for blockSize in range(MIN_BLOCK_SIZE, remainingSpaces - position + 1):
            # always at least 1 combination (this block and no others)
            total_combinations += 1
            

            next_position = position + blockSize + 1
            total_combinations += solution_two(remainingSpaces - next_position)
            
    GAP_CACHE[remainingSpaces] = total_combinations
    return total_combinations

def main():
    global LENGTH
    # start with 1 for the vacuous case
    while (solutionCount := 1 + solution_two(LENGTH)) < 10**6:
        LENGTH += 1
    print(LENGTH)

if __name__ == "__main__":
    main()
