from helpers.read_input import *

def is_duplicate(tally, num, i):
    if len(tally[num]) == 1:
        return False

    return True

def memory(list):
    tally = {}
    for i, num in enumerate(list):
        tally[num] = [i+1]

    turn = len(tally)+1
    while turn < 2021:
        lastNum = list[-1]
        # consider the last number spoken
        if is_duplicate(tally, lastNum, turn) is False:
            num = 0
            # the next one is zero
            # list.append(0)
            # tally[0].append(turn)
        else:
            # difference between last spoken and last time before that
            # get last time set
            num = tally[lastNum][-1] - tally[lastNum][-2]

            # tally[diff].append(turn)
            # list.append(diff)
        list.append(num)
        if num in tally.keys():
            tally[num].append(turn)
        else:
            tally[num] = [turn]
        turn+=1




    return list[-1]


def run():
    tests = [
       [436, [0, 3, 6]],
        [1, [1, 3, 2]],
        [10, [2, 1, 3]],
        [27, [1, 2, 3]],
        [78, [2, 3, 1]],
        [438, [3, 2, 1]],
        [1836, [3, 1, 2]],
    ]

    for t in tests:
        assert(memory(t[1]) == t[0])

    print("done")

    print(memory([6, 13, 1, 15, 2, 0]))
