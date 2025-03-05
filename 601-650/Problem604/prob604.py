import math

MAX_BOUNDS = 100

def isConvex(points: list) -> bool:
    # print(points)
    lastGradient = 0
    lastPoint = points[0]
    for p in points[1:]:
        gradient = (p[1]-lastPoint[1])/(p[0]-lastPoint[0])
        # print(gradient, lastGradient)
        if gradient <= lastGradient:
            # print("NOT CONVEX")
            return False
        lastGradient = gradient
        lastPoint = p
    # print("CONVEX")
    return True

def isNewPointConvex(points: list, candidatePoint: tuple) -> bool:
    newGradient = (candidatePoint[1]-points[-1][1])/(candidatePoint[0]-points[-1][0])
    return newGradient > getLastGradient(points)

def getLastGradient(points: list) -> float:
    if len(points) < 2:
        return 0
    return (points[-1][1]-points[-2][1])/(points[-1][0]-points[-2][0])

def getPossiblePoints(maxDimension: int, col: int, row: int, pointList: list) -> list:
    """
    Recursively gets a list of convex points from a subset of the lattice grid

    Parameters:
    @maxDimension : Defines the side length of the NxN square we are examining
    @col : Defines the column index of the bottom-left point in the lattice grid
    @row : Defines the row index of the bottom-left point in the lattice grid
    @pointList : Collection of the points in the convex curve that satisfies the constraints of the problem

    Returns:
    List of list of points satisfying the constraint of the problem
    """

    if col >= maxDimension or row >= maxDimension:
        # print(f'Returning List: {pointList}')
        return [pointList]

    listOfPointLists = []

    # For the row, the step should be ceil of previous gradient + 1
    stepAmount = math.floor(getLastGradient(pointList)) + 1

    for colIndex in range(col+1, maxDimension+1):
        for rowIndex in range(row+1, maxDimension+1):
            candidatePoint = (colIndex, rowIndex)
            # print(f'Checking ({candidatePoint[0]}, {candidatePoint[1]})')

            # Create a copy of the point list with the candidate point, and check if it's convex
            # If so, recurse with the candidate point
            # print(newPointList)

            # newPointList = pointList + [candidatePoint]
            # if isConvex(newPointList):
            if isNewPointConvex(pointList, candidatePoint):
                newPointList = pointList + [candidatePoint]
                newLists = getPossiblePoints(
                        maxDimension, 
                        candidatePoint[0], 
                        candidatePoint[1], 
                        newPointList)
                
                # print(F'New Lists: {newLists}')
                
                listOfPointLists.extend(newLists)

    return listOfPointLists

def main():
    for i in range(1, MAX_BOUNDS+1):
        pointLists = getPossiblePoints(i, 0, 0, [(0, 0)])
        maxIndex = 0
        maxLength = 0

        for j in range(len(pointLists)):
            listLength = len(pointLists[j])
            if listLength > maxLength:
                maxIndex = j
                maxLength = listLength

        print(f'F({i}) = {maxLength} - {pointLists[maxIndex]}')

    # print(isConvex([(0,0), (1,1)]))

if __name__ == "__main__":
    main()
    # print(getLastGradient([(0, 0), (2, 1), (3, 2)]))
