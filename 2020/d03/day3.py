from helpers.read_input import *

def tobaggan(forest, lat, lon):
    for l in forest:
        print(l)
    # if at any point if the right is > width, subtract width
    num_trees = 0
    right = 0
    down = 0
    
    while True:
        # are we at the bottom?
        if down >= len(forest):
            break
        # print(f"r {right} d {down} loc: {forest[down][right]}")

        # did we hit a tree?
        if forest[down][right] == "#":
            # print("hit!")
            # forest[down][right] = "X"
            num_trees+=1
        # else:
            # print("clear")
            # forest[down][right] = "O"

        # continue
        # move 3 to the right
        right += lat
        down += lon
        # are we wider than the width of the forest?
        if right >= len(forest[0]):
            right = right - len(forest[0])

    # for l in forest:
    #     print(l)

    return num_trees

def make_map(lines):
    forest = []
    for l in lines:
        latitude = []
        for char in l:
            latitude.append(char)
        forest.append(latitude)
    return forest

def run():
    lines = read_line_input("d03")
    map = make_map(lines)
    # Right 1, down 1.
    # Right 3, down 1. (This is the slope you already checked.)
    # Right 5, down 1.
    # Right 7, down 1.
    # Right 1, down 2.

    print("Num of trees hit: ", tobaggan(map, 3, 1))

    a = tobaggan(map, 1, 1)
    b = tobaggan(map, 3, 1)
    c = tobaggan(map, 5, 1)
    d = tobaggan(map, 7, 1)
    e = tobaggan(map, 1, 2)

    print("Num of trees hit: ", a, b, c, d, e)

    print(a*b*c*d*e)


run()
