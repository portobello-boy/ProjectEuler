"""
Problem 11

In the 20×20 grid below, four numbers along a diagonal line have been marked in red.

    [see grid.txt]

The product of these numbers is 26 × 63 × 78 × 14 = 1788696.

What is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20×20 grid?
"""
from MathLib.grid import fileToGrid, largestAdjacentProduct

def main():
    grid = fileToGrid("./grid.txt")
    print(largestAdjacentProduct(grid, 4))

if __name__ == "__main__":
    main()