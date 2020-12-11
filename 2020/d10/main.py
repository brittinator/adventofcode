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
            three_diff+=1
        else:
            print('something went wrong')
        prev = num

    print(ones_diff, three_diff)

    return ones_diff * three_diff


def distinct_ways(adapters):
    adapters.insert(0, 0)
    adapters.append(adapters[-1]+3)
    solutions = []


    for i, num in enumerate(adapters):
        paths = 0
        if i == 0:
            paths = 1
        if paths == 0:
            # check what previous values are

            if i-1 >=0:
                soln1 = solutions[i-1]
                if num-soln1[0] <= 3:
                    paths+=soln1[1]
            if i-2>=0:
                soln2 = solutions[i-2]
                if num-soln2[0] <= 3:
                    paths+=soln2[1]
            if i-3>=0:
                soln3 = solutions[i-3]
                if num-soln3[0] <= 3:
                    paths += soln3[1]
            
        solutions.append([num, paths])

    return solutions[-1][1]


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

    adapters = []
    for l in test_lines:
        adapters.append(l)

    adapters.sort()
    ways = distinct_ways(adapters)
    print(ways)

    assert(ways == 8)

    test_lines = read_line_input_int('d10', "test2")
    answer2 = test(test_lines)
    assert(answer2 == 220)

    adapters = []
    for l in test_lines:
        adapters.append(l)

    adapters.sort()
    ways = distinct_ways(adapters)
    print(ways)

    assert(ways == 19208)


    print("DONE")

    lines = read_line_input_int("d10")
    adapters = []
    for l in lines:
        adapters.append(l)

    adapters.sort()
    print(adapters)

    print(adapter_array(adapters))

    ways = distinct_ways(adapters)
    print(ways)
