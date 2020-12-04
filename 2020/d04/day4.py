from helpers.read_input import *


def valid_byr(byr):
    return len(byr) == 4 and 1920 <= int(byr) and int(byr) <= 2002

def valid_iyr(yr):
    return len(yr) == 4 and 2010 <= int(yr) and int(yr) <= 2020

def valid_eyr(yr):
    return len(yr) == 4 and 2020 <= int(yr) and int(yr) <= 2030

def valid_hgt(h):
    if 'cm' not in h and 'in' not in h:
        return False
    if h[len(h)-2:] == 'cm':
        h = int(h[:-2])
        return 150 <= h and h <= 193 
    elif h[len(h)-2:] == 'in':
        h = int(h[:-2])
        return 59 <= h and h <= 76

def valid_hcl(c):
    if c[0] !=  "#":
        return False
    if len(c) != 7:
        return False
    return c[1:].isalnum()

def valid_ecl(color):
    valid = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
    return color in valid

def valid_id(id):
    if len(id) != 9:
        return False
    return id.isnumeric()


def is_valid_passport(passport):
    required = {'byr': valid_byr, 'iyr': valid_iyr, 'eyr': valid_eyr,
                'hgt': valid_hgt, 'hcl': valid_hcl, 'ecl': valid_ecl, 'pid': valid_id}

    marked_fields = {}
    for k in required:
        marked_fields[k] = False

    for l in passport:
        fields = l.strip().split(" ")
        for field in fields:
            for req in required:
                if req in field:
                    _, val = field.split(f"{req}:")
                    if required[req](val) is False:
                        return False
                    marked_fields[req] = True

    vals = marked_fields.values()
    for v in vals:
        if v == False:
            return False

    return True


def valid_passports(lines):
    num_valid = 0
    pp = []
    for l in lines:
        if len(l) == 0:
            if is_valid_passport(pp):
                num_valid+=1
            pp = []
        else:
            pp.append(l)

    # test out the last passport
    if is_valid_passport(pp):
        num_valid+=1

    return num_valid
    

def run():
    lines = read_line_input("d04")
    print(valid_passports(lines))


run()
