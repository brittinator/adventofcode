from helpers.read_input import *
import re

def find_bag_of_holding(rules, bag_color):
    outer_bags = []
    for rule, bags in rules.items():
        if bag_color in bags:
            # print("match",rule,  bags)
            outer_bags.append(rule)

    return outer_bags

def create_rules(lines):
    # key is the outer bag, values are a set of inner bags
    rules = {}

    for l in lines:
        inner_bag, outer_bags = parse_line(l)
        rules[inner_bag] = outer_bags
    # print(rules)

    return rules

def parse_line(line):
    outer, inner = line.split("contain")
    outer_name, _= outer.split(" bags ")

    bags = inner.split(", ")
    inner_bags = []
    for bag in bags:
        bag = re.sub('bag[s]?', '', bag).strip()
        bag = re.sub('[0-9] ', '', bag)
        bag = bag.strip('.').strip()
        # print("inner: ", bag)
        if bag != "no other":
            inner_bags.append(bag)

    return outer_name, inner_bags


def runner(rules, color, already_counted):
    matches = find_bag_of_holding(rules, color)

    for color in matches:
        if color not in already_counted:
            already_counted.add(color)
            mc = runner(rules, color, already_counted)
            print("MC", mc)
            # if mc is not None:
                # match_count += mc

    return already_counted



def test(lines):
    rules = create_rules(lines)
    find_bag_of_holding(rules, 'shiny gold')
    parse_line("light red bags contain 1 bright white bag, 2 muted yellow bags.")

    a = set()
    mc = runner(rules, 'shiny gold', a)

    print("matches: ", len(mc), mc, a)

def run():
    # test_lines = read_line_input('d07', "test")
    # print(test(test_lines))
    # test(test_lines)


    lines = read_line_input("d07")
    rules = create_rules(lines)

    # match_count = 0
    matches = set()
    runner(rules, 'shiny gold', matches)
    print("num matches: ", len(matches))

    # print("matches: ", match_count)


run()
