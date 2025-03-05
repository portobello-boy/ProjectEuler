LENGTH = 8
MIN_BLOCK_SIZE = 3
PARTITION = [LENGTH]

GAP_CACHE = {}
CHECKED_INDICES = set()

def solution(startPosition, endPosition):
    if endPosition - startPosition < MIN_BLOCK_SIZE:
        return 0
    
    if (startPosition, endPosition) in CHECKED_INDICES:
        return 0
    CHECKED_INDICES.add((startPosition, endPosition))
    
    
    print(f'recursing: {startPosition} - {endPosition}')
    
    if endPosition - startPosition in GAP_CACHE:
        return GAP_CACHE[endPosition-startPosition]
    
    
    totalCombinations = 0
    # iterate over block sizes valid for range
    for blockSize in range(MIN_BLOCK_SIZE, endPosition - startPosition + 1):
        # iterate over positions in this range
        for position in range(startPosition, endPosition - blockSize + 1):
            
            print(f''.rjust(position) + f''.rjust(blockSize, '-'))

            # recurse for left and right
            leftSolution = solution(startPosition, position - 1)
            rightSolution = solution(position + blockSize + 1, endPosition)
            
            totalCombinations += 1 + leftSolution + rightSolution

            # if leftSolution > 0:
            #     print(f'{startPosition}'.rjust(startPosition) + f' {leftSolution} ' + f'{position - 1}')
            # if rightSolution > 0:
            #     print(f'{position + blockSize + 1}'.rjust(position + blockSize + 1) + f' {rightSolution} ' + f'{endPosition}')

    GAP_CACHE[endPosition-startPosition] = totalCombinations
    print('returning')
    return totalCombinations
    

def main():
    print(solution(0, LENGTH))
    print(GAP_CACHE)

if __name__ == "__main__":
    main()
