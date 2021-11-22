"""
Operations on grids of numbers
"""

from typing import List
from more_itertools import windowed
from functools import reduce

def fileToGrid(file:str) -> List[List[int]]:
    grid = list()
    with open(file, "r") as readFile:
        for row in readFile:
            digits = [int(d) for d in row.split()]
            # print(digits)
            grid.append(digits)
    return grid

def transposeGrid(grid:List[List[int]]) -> List[List[int]]:
    return [[grid[j][i] for j in range(len(grid))] for i in range(len(grid[0]))]

def largestHorizontalProduct(row:List[int], size:int) -> int:
    maxProd = None
    for w in windowed(row, size):
        if not maxProd:
            maxProd = reduce(lambda x, y: x * y, w)
        else:
            maxProd = max(maxProd, reduce(lambda x, y: x * y, w))
    return maxProd

def largestDiagonalProduct(grid:List[List[int]], size:int) -> int:
    maxProd = None

    # Positive diagonal
    for i in range(len(grid) - size):
        for j in range(len(grid[0]) - size):
            w = list()
            for k in range(size):
                w.append(grid[i+k][j+k])
            if not maxProd:
                maxProd = reduce(lambda x, y: x * y, w)
            else:
                maxProd = max(maxProd, reduce(lambda x, y: x * y, w))
    
    # Negative diagonal
    for i in range(len(grid) - size):
        for j in range(size-1, len(grid[0])):
            w = list()
            for k in range(size):
                w.append(grid[i+k][j-k])
            maxProd = max(maxProd, reduce(lambda x, y: x * y, w))
    
    return maxProd            

def largestAdjacentProduct(grid:List[List[int]], size:int) -> int:
    maxProd = None
    for r in grid:
        if not maxProd:
            maxProd = largestHorizontalProduct(r, size)
        else:
            maxProd = max(maxProd, largestHorizontalProduct(r, size))

    for r in transposeGrid(grid):
        maxProd = max(maxProd, largestHorizontalProduct(r, size))

    return max(maxProd, largestDiagonalProduct(grid, size))
