from helpers.read_input import *

# row 
# 0-127
# F = lower (0-63)
# B = higher upper (64-127)

# col
# 0-7
# L lower left
# R upper right

# final thing: row *8 + col

def row_search(input, lo, hi):
    if lo == hi or len(input) == 0:
        return lo

    for i, char in enumerate(input):
        mid = (hi+lo)//2

        if char == "L":
            return row_search(input[i+1:], lo, mid-1)
        elif char == "R":
            return row_search(input[i+1:], mid+1, hi)
        else:
            exit()
    print("at end, returning with nothing")


def col_search(input, lo, hi):
    if lo == hi or len(input) == 0:
        return lo

    for i, char in enumerate(input):
        mid = (hi+lo)//2

        if char == "F":
            return col_search(input[i+1:], lo, mid-1)

        elif char == "B":
            return col_search(input[i+1:], mid+1, hi)
        else:
            exit()

    print("at end, returning with nothing")


def boarding_seat(bpass):
    col = col_search(bpass[0:7], 0, 127)

    row = row_search(bpass[-3:], 0, 7)
    # print("row is: ", row)

    return col*8 + row

def find_seat(passes):
    seats = []

    for bpass in passes:
        seat = boarding_seat(bpass)
        seats.append(seat)
    seats.sort()

    print('seats ',seats)
    supposed_seats = []
    min = seats[0]
    max = seats[-1]
    for i in range(min, max+1):
        supposed_seats.append(i)

    print("SUPPOSED", len(supposed_seats))

    print("HERE")
    # return set(range(min, max+1)).difference(seats)

    # print("SEATS: ", len(seats))
    # prev = seats[0]
    for i in range(min, max+1):
        if i not in seats:
            return i


def binary_boarding(passes):
    highest_seat = 0
    for bpass in passes:
        seat = boarding_seat(bpass)
        print(seat, highest_seat)
        if seat > highest_seat:
            highest_seat = seat
    
    return highest_seat


def test():
    assert col_search("BFB", 0, 4) == 3
    test_input = "bfffbbfrrr".upper()
    assert col_search(test_input, 0, 127) == 70

    assert row_search("RRR", 0, 7) ==  7

    result = boarding_seat(test_input)
    print("result: ", result)
    assert(result) == 567

    assert(boarding_seat("BBFFBBFRLL")) == 820

    print("assert DONE")

def run():
    test()
    
    lines = read_line_input("d05")
    print("highest: ", binary_boarding(lines))

    # print(find_seat(read_line_input("d05")))
