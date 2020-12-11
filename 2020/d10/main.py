from helpers.read_input import *

def adapter_array(nums):
    nums.append(nums[-1]+3)
    ones_diff = 0
    three_diff = 0
    prev = 0
    for num in nums:
        diff = num-prev
        if diff == 1:
            ones_diff+=1
        elif diff == 3:
            print("three diff: ", num)
            three_diff+=1
        else:
            print('something went wrong')
        prev = num

    print(ones_diff, three_diff)

    return ones_diff * three_diff


def test(lines):
    adapters = []
    for l in lines:
        adapters.append(l)

    adapters.sort()
    print(adapters)


    return adapter_array(adapters)



def run():
    test_lines = read_line_input_int('d10', "test")
    answer1 = test(test_lines)
    assert(answer1 == 35)

    test_lines = read_line_input_int('d10', "test2")
    answer2 = test(test_lines)
    assert(answer2 == 220)

    print("DONE")

    lines = read_line_input_int("d10")
    adapters = []
    for l in lines:
        adapters.append(l)

    adapters.sort()
    print(adapters)

    print(adapter_array(adapters))
