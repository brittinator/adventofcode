from helpers.read_input import *
from collections import deque

def has_match(preamble, num):
    for match in preamble:
        diff = num-match
        if diff == num or diff < 0:
            continue

        # does this exist in the set?
        for n in preamble:
            if diff == n:
                return True
        # if diff in preamble == True:
        #     return True

    return False

def new_q(items, maxlen):
    q = deque(maxlen=maxlen)
    for item in items:
        q.append(item)

    return q

def runner(lines, size):
    items = []
    # create q
    i = 0
    while i < size:
        items.append(int(lines[i]))
        i += 1

    q = new_q(items, size)

    print("i is: ", i)
    while i < len(lines):
        num = int(lines[i])
        # print("i num", i, num)
        if has_match(q, num) is False:
            return num

        q.append(num)
        i += 1


def test(lines):
   return runner(lines, 5)

def run():
    test_lines = read_line_input('d09', "test")
    print(test(test_lines))

    lines = read_line_input("d09")
    print(runner(lines, 25))
