from read_input import *

def accounting():
    nums = read_line_input_int("01")
    matches = {}
    pairs = []
    for n in nums:
        pair = 2020-n
        if matches.get(pair) is not None:
            print("Matched!")
            pairs = [n, pair]
            break
        matches[n] = 0

    print(pairs)
    print(pairs[0] * pairs[1])


accounting()
    
