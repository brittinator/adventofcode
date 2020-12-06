
from helpers.read_input import *


def get_all_count(unique, group_count):
    count = 0
    # get the key with a value of group_length
    for _, v in unique.items():
        if v == group_count:
            count += 1

    return count


def customs(lines):
    counts = []
    everyone_counts = []

    unique = {}
    group_count = 0
    for answers in lines:
        if len(answers) == 0 :
            # end of group
           #  print("U", len(unique), unique, group_count)
            
            everyone_counts.append(get_all_count(unique, group_count))
            counts.append(len(unique))

            unique = {}
            group_count = 0
        else:
            group_count+=1
        for answer in answers:
            unique[answer] = unique.get(answer, 0) +1

    everyone_counts.append(get_all_count(unique, group_count))
    counts.append(len(unique))


    return sum(counts), sum(everyone_counts)




def test():
    assert(True is True)

    
    print("ASSERT done")

def run():
    test()

    lines = read_line_input("d06")

    print(customs(lines))
