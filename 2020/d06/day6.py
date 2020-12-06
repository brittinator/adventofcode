from helpers.read_input import *

def get_all_count(unique, group_count):
    all_count = filter(lambda v: v == group_count, unique.values())
    
    return len(list(all_count))

def customs(lines):
    counts = []
    everyone_has_these_counts = []

    unique = {}
    group_count = 0
    for answers in lines:
        if len(answers) == 0 :
            everyone_has_these_counts.append(get_all_count(unique, group_count))
            counts.append(len(unique))

            unique = {}
            group_count = 0
        else:
            group_count+=1
        for answer in answers:
            unique[answer] = unique.get(answer, 0) +1

    everyone_has_these_counts.append(get_all_count(unique, group_count))
    counts.append(len(unique))

    return sum(counts), sum(everyone_has_these_counts)

def run():
    lines = read_line_input("d06")

    print(customs(lines))
