from helpers.read_input import *

def search(input, lo, hi):
    # print(input, lo, hi)
    mid = (hi+lo)//2
    if len(input) == 0:
        return lo

    for i, char in enumerate(input):
        if char == "L" or char =="F":
            return search(input[i+1:], lo, mid)
        else:
            return search(input[i+1:], mid+1, hi)

def boarding_seat(bpass):
    col = search(bpass[0:7], 0, 127)
    row = search(bpass[-3:], 0, 7)

    return col*8 + row

def find_seat(seats):
    return set(range(seats[0], seats[-1])).difference(seats)

def binary_boarding(passes):
    seats = []
    highest_seat = 0
    for bpass in passes:
        seat = boarding_seat(bpass)
        seats.append(seat)
        if seat > highest_seat:
            highest_seat = seat
    
    return seats, highest_seat

def test():
    test_input = "bfffbbfrrr".upper()
    assert search(test_input[:-3], 0, 127) == 70

    assert (search("RRR", 0, 7) == 7), search("RRR", 0, 7)

    result = boarding_seat(test_input)
    assert(result) == 567

    assert(boarding_seat("BBFFBBFRLL")) == 820

    print("assert DONE")

def run():
    test()
    
    lines = read_line_input("d05")
    seats, bb = binary_boarding(lines)
    print("highest: ", bb)

    # print(sorted(seats))

    print(find_seat(sorted(seats)))
