from types import FunctionType
from typing import List


def generatePolynomial(coeffs:List[float]) -> FunctionType:
    def f(x:float) -> float:
        total = 0
        for i, c in enumerate(reversed(coeffs)):
            total += c * x**i
        return total
    return f