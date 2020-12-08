from helpers.read_input import *
import copy

def parse_input(lines):
    # idx is the location, val is if it's been visited
    visited = []
    for i, l in enumerate(lines):
        instruct, arg = l.split(' ')
        instructions = {
            'instruct':instruct,
             'arg': int(arg),
             'seen': False,
        }
        visited.append(instructions)

    return visited

def is_repeated(instructions, accum):
    repeats = True
    i = 0
    while i < len(instructions) and repeats == True:
        new_instructions = copy.deepcopy(instructions)
        instruct = new_instructions[i]['instruct']
        print(i, instruct)
        accum = 0
        if instruct == 'jmp':
            new_instructions[i]['instruct'] = 'nop'
        elif instruct == 'nop':
            new_instructions[i]['instruct'] = 'jmp'
        else:
            i+=1
            continue

        accum, repeats= loop_accumulator(new_instructions, 0)
        i+=1

    return accum

def loop_accumulator(visited, accum):
    # start at the beginning
    i = 0

    while i < len(visited) and visited[i]['seen'] == False:
        instruct = visited[i]['instruct']
        # default move of 1
        move = 1
        if instruct == 'acc':
            accum+= visited[i]['arg']
        elif instruct == 'jmp':
            move= visited[i]['arg']
        visited[i]['seen'] = True
        i+=move
    if i >= len(visited):
        return accum, False

    return accum, visited[i]['seen']

def test():
    lines = read_line_input("d08", 'test')
    instructions = parse_input(lines)
    
    accum = is_repeated(instructions, 0)

    assert(accum == 8)
    print("done with assert", accum)


def run():
    test()

    lines = read_line_input("d08")
    instructions = parse_input(lines)
    # print(loop_accumulator(instructions, 0))
    print(is_repeated(instructions, 0))
