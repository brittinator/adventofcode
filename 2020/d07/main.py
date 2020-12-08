from helpers.read_input import *
import re

def find_bag_of_holding(rules, bag_color):
    outer_bags = []
    for rule, bags in rules.items():
        if bag_color in bags:
            outer_bags.append(rule)

    return outer_bags

def create_rules(lines):
    # key is the outer bag, values are a set of inner bags
    rules = {}

    for l in lines:
        outer_bag, inner_bags = parse_line(l)
        rules[outer_bag] = inner_bags

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
        if bag != "no other":
            inner_bags.append(bag)

    return outer_name, inner_bags


def create_rules_part_2(lines):
    # key is the outer bag, values are a set of inner bags
    rules = {}

    for l in lines:
        outer_bag, inner_bags = parse_line_part_2(l)
        rules[outer_bag] = inner_bags

    return rules

def parse_line_part_2(line):
    outer, inner = line.split("contain")
    outer_name, _ = outer.split(" bags ")

    bags = inner.split(", ")
    inner_bags = {}
    for bag in bags:
        bag = re.sub('bag[s]?', '', bag.strip()).strip()
        count = bag[0]
        bag = re.sub('[0-9] ', '', bag)
        bag = bag.strip('.').strip()
        if bag != "no other":
            inner_bags[bag] = int(count)

    return outer_name, inner_bags


def how_many_bags_inside(color, rules, chosen_colors):
    count = 0
    bags = rules[color]
    if len(bags) == 0:
        return 0

    for color, bag_count in bags.items():
        if color in chosen_colors:
            # saved result
            result = chosen_colors[color]
        else:
            # get the contents
            result = how_many_bags_inside(color, rules, chosen_colors)
            # add it to chosen
            chosen_colors[color] = result
        count+= bag_count
        count += bag_count*result


    return count


def runner(rules, color, already_counted):
    matches = find_bag_of_holding(rules, color)

    for color in matches:
        if color not in already_counted:
            already_counted.add(color)
            runner(rules, color, already_counted)

    return already_counted

def runner_part_2():
    lines = read_line_input("d07")
    rules = create_rules_part_2(lines)

    chosen_colors = {}
    print("how many?", how_many_bags_inside('shiny gold', rules, chosen_colors))


def test(lines):
    rules = create_rules(lines)
    find_bag_of_holding(rules, 'shiny gold')
    parse_line("light red bags contain 1 bright white bag, 2 muted yellow bags.")

    a = set()
    mc = runner(rules, 'shiny gold', a)

    print("matches: ", len(mc), mc, a)

def run():
    test_lines = read_line_input('d07', "test")
    # print(test(test_lines))
    test(test_lines)


    lines = read_line_input("d07")
    rules = create_rules(lines)

    matches = set()
    runner(rules, 'shiny gold', matches)
    print("num matches: ", len(matches))

    runner_part_2()

