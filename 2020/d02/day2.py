from helpers.read_input import *

def num_of_valid_pw(lines):
    num_valid_pw = 0
    for l in lines:
        key, pw = l.split(":")
        pw = pw.strip()
        range, char = key.split(" ")
        lo, hi = range.split("-")
        lo = int(lo.strip())
        hi = int(hi.strip())

        if is_valid_for_realz(pw, lo, hi, char):
            num_valid_pw+=1

    return num_valid_pw

def is_valid(pw, lo, hi, special_char):
    char_count = 0
    for char in pw:
        if char == special_char:
            char_count+=1
    
    return char_count >= int(lo) and char_count <= int(hi)

def is_valid_for_realz(pw, loc1, loc2, special_char):
    # print(loc1, loc2, pw, special_char)
    l1_bool = pw[loc1-1] == special_char
    l2_bool = pw[loc2-1] == special_char

    if l1_bool is True and l2_bool is True:
        return False
    elif l1_bool is True or l2_bool is True:
        return True 

    return False

def run():
    lines = read_line_input("d02")
    print(num_of_valid_pw(lines))
