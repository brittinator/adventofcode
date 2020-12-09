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

    return False

def new_q(items, maxlen):
    q = deque(maxlen=maxlen)
    for item in items:
        q.append(item)

    return q

def runner_pt_1(lines, size):
    items = []
    # create q
    i = 0
    while i < size:
        items.append(int(lines[i]))
        i += 1

    q = new_q(items, size)

    while i < len(lines):
        num = int(lines[i])
        if has_match(q, num) is False:
            return num, i

        q.append(num)
        i += 1

def runner_pt_2(lines, invalid_num, invalid_idx):
# work backwards from the invalid number, to see if adding up numbers before equal the invalid num.
# when they don't, subtract the highest indexed one and add to the lower side until you go over.
    i = invalid_idx -1
    while i > 0:
        i_top = i
        i_bottom = i
        count = 0
        while count < invalid_num:
            count+=lines[i_bottom]
            i_bottom-=1
        if count == invalid_num:
            # find largest & smallest
            smallest = lines[i_bottom]
            largest = lines[i_bottom]
            while i <= i_top and i > i_bottom:
                curr = lines[i]
                if curr < smallest:
                    smallest = curr
                elif curr > largest:
                    largest = curr
                i-=1
            return largest, smallest 
        
        i-=1



def test(lines):
    num, invalid_idx =  runner_pt_1(lines, 5)

    for i, l in enumerate(lines):
        lines[i] = int(l)

    largest, smallest = runner_pt_2(lines, num, invalid_idx)
    print(largest, smallest)
    print("+", largest+smallest)
    

def run():
    test_lines = read_line_input('d09', "test")
    print(test(test_lines))

    lines = read_line_input("d09")
    for i, l in enumerate(lines):
        lines[i] = int(l)

    num, index = runner_pt_1(lines, 25)
    print(num)
    largest, smallest = runner_pt_2(lines, num, index)
    print(largest, smallest, largest+smallest)


