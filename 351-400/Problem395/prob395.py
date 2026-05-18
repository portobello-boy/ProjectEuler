import numpy as np
from math import cos, sin, asin
from queue import Queue
from p5 import *

TRIANGLE_SHORT_SIDE = 3
TRIANGLE_LONG_SIDE = 4
TRIANGLE_HYPOTENUSE = 5

theta = asin(TRIANGLE_SHORT_SIDE/TRIANGLE_HYPOTENUSE) # Angle of rotation
R = np.matrix([[cos(theta),-sin(theta)],
               [sin(theta), cos(theta)]]) # Rotation affine transformation matrix
Sl = np.matrix([[TRIANGLE_LONG_SIDE/TRIANGLE_HYPOTENUSE, 0],
               [0,                                       TRIANGLE_LONG_SIDE/TRIANGLE_HYPOTENUSE]]) # Matrix to scale larger child squares
Sr = np.matrix([[TRIANGLE_SHORT_SIDE/TRIANGLE_HYPOTENUSE, 0],
               [0,                                        TRIANGLE_SHORT_SIDE/TRIANGLE_HYPOTENUSE]]) # Matrix to scale smaller child squares

start_dimension = 1
depth = 25
min_left, min_right, min_bottom, min_top = 0,start_dimension,start_dimension*10,start_dimension

def find_left_child(current):
    return (Sl@R@(current-current.T[0].T))+current.T[1].T

def find_right_child(current):
    tmp = ((Sr@R@(current-current.T[0].T))+current.T[3].T).T
    return np.concatenate((tmp[1], tmp[3], tmp[0], tmp[2])).T

def find_side_length(square):
    tmp = square.T
    return np.linalg.norm(tmp[0]-tmp[1])

def pushes_boundary(square):
    global min_left, min_right, min_bottom, min_top
    pushes = False
    for p in square.T:
        if p[0,0] <= min_left:
            min_left = p[0,0]
            pushes = True
        if p[0,0] >= min_right:
            min_right = p[0,0]
            pushes = True
        if p[0,1] <= min_bottom:
            min_bottom = p[0,1]
            pushes = True
        if p[0,1] >= min_top:
            min_top = p[0,1]
            pushes = True
    return pushes

# Initialize starting square
start = np.matrix([[0,0,start_dimension,start_dimension],
                  [0,start_dimension,0,start_dimension]])

# Initialize queue and generation history of the squares (for drawing)
generation_history = {}
square_queue = Queue()
square_queue.put(start)

# Iterate up to the specified depth to calculate rectangle area
for i in range(depth+1):
    generation = {}
    generation["top"] = min_top
    generation["bottom"] = min(min_bottom, 0)
    generation["left"] = min_left
    generation["right"] = min_right
    generation_history[i] = generation
    generation["squares"] = []
    
    # Create a new queue
    new_queue = Queue()
    
    # Iterate over all leaves of the tree
    for j in range(square_queue.qsize()):
        sq = square_queue.get()

        # Find left and right child squares
        left_child = find_left_child(sq)
        right_child = find_right_child(sq)
        
        # If either pushes the maximum area, add that to the search
        if pushes_boundary(left_child):
            new_queue.put(left_child)
        if pushes_boundary(right_child):
            new_queue.put(right_child)
    square_queue = new_queue

# Iterate over all squares for drawing
square_queue = Queue()
square_queue.put(start)
for i in range(depth+1):
    new_queue = Queue()
    generation_list = []
    for j in range(square_queue.qsize()):
        sq = square_queue.get()
        side_length = find_side_length(sq)
        if side_length<0.015:
            continue
        generation_list.append(sq)

        left_child = find_left_child(sq)
        right_child = find_right_child(sq)
        new_queue.put(left_child)
        new_queue.put(right_child)
    generation_history[i]["squares"] = generation_list
    square_queue = new_queue

print(min_left, min_right, min_bottom, min_top)
print("BOTTOM LENGTH: {}".format(abs(min_bottom-min_top)))
print("SIDE LENGTH: {}".format(abs(min_right-min_left)))
print("AREA: {}".format(abs(min_right-min_left)*abs(min_bottom-min_top)))

SCALE = 75
width=1600
height=900

def setup():
    size(width,height)
    background(255)

def unit_vector(vector):
    """ Returns the unit vector of the vector.  """
    return vector / np.linalg.norm(vector)

def angle_between(v1, v2):
    """ Returns the angle in radians between vectors 'v1' and 'v2'::

            >>> angle_between((1, 0, 0), (0, 1, 0))
            1.5707963267948966
            >>> angle_between((1, 0, 0), (1, 0, 0))
            0.0
            >>> angle_between((1, 0, 0), (-1, 0, 0))
            3.141592653589793
    """
    v1_u = unit_vector(v1)
    v2_u = unit_vector(v2)
    return np.arccos(np.dot(v1_u, v2_u))

prev_squares=[]
generation_iterator = 0

tree_fill_color = (107, 71,  0, 200)
tree_edge_color = (107, 71,  0, 255)
leaf_fill_color = (0,   102, 0, 200)
leaf_edge_color = (0,   102, 0, 255)

def draw():
    global generation_iterator
    background(255)
    set_frame_rate(1)
    translate(width/2, 2*height/3)
    rotate(PI)
    
    generation = generation_history[generation_iterator]
    
    line(SCALE*generation["left"], SCALE*generation["bottom"], SCALE*generation["right"], SCALE*generation["bottom"])
    line(SCALE*generation["left"], SCALE*generation["bottom"], SCALE*generation["left"], SCALE*generation["top"])
    line(SCALE*generation["left"], SCALE*generation["top"], SCALE*generation["right"], SCALE*generation["top"])
    line(SCALE*generation["right"], SCALE*generation["bottom"], SCALE*generation["right"], SCALE*generation["top"])
    
    for sq in prev_squares:
        tmp = sq.T
        rotation_angle = angle_between(tmp[1]-tmp[0], np.matrix([[1], [0]]))[0,0]
        if (tmp[1]-tmp[0])[0,1] < 0:
            rotation_angle *= -1
            
        push()
        translate(tmp[0,0]*SCALE, tmp[0,1]*SCALE)
        stroke_weight(1)
        rotate(rotation_angle - PI/2)
        stroke_weight(1)
        stroke(*tree_edge_color)
        fill(*tree_fill_color)
        square(0,0, find_side_length(sq)*SCALE)
        pop()
    
    for sq in generation["squares"]:
        prev_squares.append(sq)
        tmp = sq.T
        rotation_angle = angle_between(tmp[1]-tmp[0], np.matrix([[1], [0]]))[0,0]
        if (tmp[1]-tmp[0])[0,1] < 0:
            rotation_angle *= -1
            
        push()
        translate(tmp[0,0]*SCALE, tmp[0,1]*SCALE)
        stroke_weight(1)
        rotate(rotation_angle - PI/2)
        stroke_weight(1)
        stroke(*leaf_edge_color)
        fill(*leaf_fill_color)
        square(0,0, find_side_length(sq)*SCALE)
        pop()
        
    generation_iterator+=1
    if generation_iterator==len(generation_history):
        no_loop()
    
run(renderer="skia")

