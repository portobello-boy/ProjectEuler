from typing import List

def magnitude(n:List[int]) -> int:
    return sum([v**2 for v in n])