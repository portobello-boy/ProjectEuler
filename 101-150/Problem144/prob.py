# Written by Cameron Haddock & Daniel Millson
# Submitted as a solution to Project Euler's Problem 144

from sympy import Symbol, solve
import math
from math import isclose

A = 5  # semi-minor axis
B = 10 # semi-major axis
INITIAL_C = B + 0.1
INITIAL_M = -19.7/1.4
PRECISION = 50


def point_on_ellipse(x, y) -> bool:
    return isclose(4*(x**2) + (y**2), 100, abs_tol=0.001)

def get_tangent_slope(x, y) -> float:
    return -4 * x / y

def get_normal_slope(x, y) -> float:
    return -1 / get_tangent_slope(x, y)

def points_of_intersect(m, c, a = A, b = B) -> (float, float):
    x = Symbol('x')
    line_eq = m*x + c
    intersect_eq = ((a**2)*(m**2) + b**2)*(x**2) + (2*(a**2)*m*c)*x + (a**2 * (c**2 - b**2))
    
    roots = solve(intersect_eq)
    
    y1 = line_eq.subs(x, roots[0])
    y2 = line_eq.subs(x, roots[1])

    p1 = (round(roots[0], PRECISION), round(y1, PRECISION))
    p2 = (round(roots[1], PRECISION), round(y2, PRECISION))

    return {p1, p2}

# http://www.sdmath.com/math/geometry/reflection_across_line.html
def reflect_point_across_line(x, y, m, c) -> (float, float):
    u = ((1 - m**2)*x + 2*m*y - 2*m*c)/((m**2) + 1)
    v = ((m**2 - 1)*y + 2*m*x + 2*c)/((m**2) + 1)
    return u, v

def points_are_close(x1, y1, x2, y2, p=4) -> bool:
    return (((x1-x2)**2) + ((y1-y2)**2))**0.5 < 10**-p

def order_points(points,old_point) -> [(float, float), (float, float)]:
    points = tuple(points)
    p1, p2 = points[0], points[1]

    if points_are_close(*p1,*old_point):
        return p2, old_point
    return p1, old_point
    

def main():
    m, c = INITIAL_M, INITIAL_C
    old_point = sorted(points_of_intersect(m, c))[0] # Extract first point of intersection of ellipse
    x, y = 0, -10
    
    i = 0
    while x < -0.01 or x > 0.01 or y < 0:
        points = points_of_intersect(m, c)
        point_of_reflection, old_point = order_points(points, old_point) #next(iter(points-{old_point}))
        
        # print(points-old_point)
        # if not point_on_ellipse(*point_of_reflection):
        #     print("POR not on ellipse")
        # print(point_of_reflection, old_point)

        normal_slope = get_normal_slope(*point_of_reflection)
        normal_c = point_of_reflection[1] - normal_slope * point_of_reflection[0]

        reflected_point = reflect_point_across_line(*old_point, normal_slope, normal_c)

        m = (point_of_reflection[1] - reflected_point[1]) / (point_of_reflection[0] - reflected_point[0])
        c = point_of_reflection[1] - m * point_of_reflection[0]

        old_point = point_of_reflection
        x, y = old_point[0], old_point[1]
        i += 1
    print(i-1)

if __name__ == "__main__":
    main()