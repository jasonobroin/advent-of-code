# https://adventofcode.com/2021/day/5
#
# Read the input from stdin

import sys

def print_array(puzzle):
    for y in range(0, max_y):
        s = ''
        for x in range(0, max_x):
            if puzzle[y][x] == 0:
                s += '.'
            else:
                s += str(puzzle[y][x])
        print(s)

def calc_overlap(puzzle):
    overlap = 0
    for y in range(0, max_y):
        for x in range(0, max_x):
            if puzzle[y][x] >= 2:
                overlap += 1
    return overlap

puzzle = []

max_x = 0
max_y = 0

for line in sys.stdin:
    coords = line.strip().split(" -> ")
    first = [int(s) for s in coords[0].split(",")]
    second = [int(s) for s in coords[1].split(",")]

    # Build array - expanding on demand

    # print(first, second)

    if first[0] > second[0] and first[1] == second[1]:
        # reverse direction
        third = second[0]
        second[0] = first[0]
        first[0] = third
    elif first[1] > second[1] and first[0] == second[0]:
        # reverse direction
        third = second[1]
        second[1] = first[1]
        first[1] = third

    # print("          -> ", first, second)

    if second[0] >= max_x or second[1] >= max_y:
        # Need to expand the array
        # Start by extending each existing line if required
        if second[0] >= max_x:
            for y in range(0, max_y):
                for e in range(0, (second[0] - max_x)):
                    puzzle[y].append(0)
            max_x = second[0]
        # Then add new lines of new length (if required)
        if second[1] >= max_y:
            for y in range(max_y, second[1] + 1):
                newline = [0] * (max_x + 1)
                puzzle.append(newline)
            max_y = second[1] + 1

    # print(f"max x {max_x} max_y {max_y}")

    if first[1] == second[1]:
        # Fill horizontally
        for x in range(first[0], second[0] + 1):
            puzzle[first[1]][x] += 1
    elif first[0] == second[0]:
        # Fill vertically
        for y in range(first[1], second[1] + 1):
            puzzle[y][first[0]] += 1
    else:
        pass

    # print_array(puzzle)

# pt1

print(calc_overlap(puzzle))
