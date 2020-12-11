from helpers.read_input import *
import copy
import unittest

# If a seat is empty(L) and there are no occupied seats adjacent to it, the seat becomes occupied.
# If a seat is occupied(  # ) and four or more seats adjacent to it are also occupied, the seat becomes empty.
# Otherwise, the seat's state does not change.

def in_bounds(plane, x,y):
    return 0<=x and x <=len(plane) -1 and y>= 0 and y <= len(plane[0])-1

def safe_get(plane, x,y):
    if in_bounds(plane, x, y) is False:
        return ''
    return plane[x][y]

def check_top(plane, x,y):
    return safe_get(plane, x-1, y-1) + safe_get(plane, x-1, y) + safe_get(plane, x-1, y+1)
    

def check_bottom(plane, x,y):
    return safe_get(plane, x+1, y-1) + safe_get(plane, x+1, y) + safe_get(plane, x+1, y+1)

def check_sides(plane, x,y):
    return safe_get(plane, x, y-1) + safe_get(plane, x, y+1)

def round(plane):
    new_plane = copy.deepcopy(plane)
    for x, row in enumerate(plane):
        for y, seat in enumerate(row):
            top = check_top(plane, x, y)
            bot = check_bottom(plane, x, y)
            mid = check_sides(plane, x, y)
            adjacents = top + bot + mid
            # if adjacents is None:
            #     continue
            empty = adjacents.count('L')
            occupied = adjacents.count('#')
            if seat == '#' and occupied >= 4:
                seat = 'L'
               
            elif seat =='L' and occupied == 0:
                seat = '#'

            new_plane[x][y] = seat

    # print("new plane")
    # for row in new_plane:
    #     print(row)
    return new_plane

def count_occupied(plane):
    count = 0
    for row in plane:
        count+=row.count('#')

    return count


def test(lines):
    round1 = round(lines)

    round1_array = []
    for row in round1:
        round1_array.append("".join(row))

    expected = ['#.##.##.##',
'#######.##',
'#.#.#..#..',
'####.##.##',
'#.##.##.##',
'#.#####.##',
'..#.#.....',
'##########',
'#.######.#',
'#.#####.##']

    assert(len(round1_array) == len(expected))

    for i, row in enumerate(expected):
        assert(row == round1_array[i], i)

    assert(round1_array == expected)
    print("DONE")


    round2 = round(round1)

    expected = [
        '#.LL.L#.##',
        '#LLLLLL.L#',
        'L.L.L..L..',
        '#LLL.LL.L#',
        '#.LL.LL.LL',
        '#.LLLL#.##',
        '..L.L.....',
        '#LLLLLLLL#',
        '#.LLLLLL.L',
        '#.#LLLL.##',
    ]

    round2_array = []
    for row in round2:
        round2_array.append("".join(row))

    for i, row in enumerate(expected):
        print(">>>>>")
        print(row)
        print(round2_array[i])
        assert(row == round2_array[i])

    assert(round2_array == expected)
    # unittest.TestCase.assertNotEqual(self, round1, round2)
    round3 = round(round2)
    round4 = round(round3)
    round5 = round(round4)
    round6 = round(round5)

    assert(round5 == round6)

    assert(count_occupied(round5) == 37)

    print("DONNNEE")


def run():
    test_lines = read_line_input('d11', "test")

    plane = []
    for row in test_lines:
        new_row = []
        for seat in row:
            new_row.append(seat)
        plane.append(new_row)

    assert(in_bounds(plane, 0,0) == True)
    assert(in_bounds(plane, 9, 10) == False)
    assert(in_bounds(plane, 9,9) == True)
    assert(in_bounds(plane, 9,0) == True)
    assert(in_bounds(plane, 10, 0) == False)


    lines = read_line_input("d11")
    plane = []
    for row in lines:
        new_row = []
        for seat in row:
            new_row.append(seat)
        plane.append(new_row)

    old_seats = plane
    new_seats = round(plane)
    while new_seats != old_seats:
        new = round(new_seats)
        old_seats = new_seats
        new_seats = new

    print("occupied: ", count_occupied(new_seats))

