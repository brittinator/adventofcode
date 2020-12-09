from typing import List

def read_line_input(day:str, file='input') -> List[str]:
    lines = []
    with open(f"{day}/{file}") as file:
        for line in file:
            lines.append(line.strip())

    return lines


def read_line_input_int(day: str) -> List[int]:
    lines = []
    with open(f"{day}/input") as file:
        for line in file:
            lines.append(int(line.strip()))

    return lines


def read_singleline(day:str) -> List[str]:
    vals = []
    with open(f"{day}/input") as file:
        for line in file:
            vals = line.split(",")

    return vals


def read_singleline_int(day: str) -> List[int]:
    vals = []
    with open(f"{day}/input") as file:
        for line in file:
            vals = line.split(",")

    for i, value in enumerate(vals):
        vals[i] = int(value)

    return vals
