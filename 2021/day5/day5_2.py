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

    x_dir = 0
    y_dir = 0
    xm = first[0]
    ym = first[1]

    if first[0] > second[0]:
        x_dir = -1
    elif first[0] < second[0]:
        x_dir = 1
        xm = second[0]

    if first[1] > second[1]:
        y_dir = -1
    elif first[1] < second[1]:
        y_dir = 1
        ym = second[1]

    if xm >= max_x or ym >= max_y:
        # Need to expand the array
        # Start by extending each existing line if required
        if xm >= max_x:
            for y in range(0, max_y):
                for e in range(0, (xm - max_x)):
                    puzzle[y].append(0)
            max_x = xm
        # Then add new lines of new length (if required)
        if ym >= max_y:
            for y in range(max_y, ym + 1):
                newline = [0] * (max_x + 1)
                puzzle.append(newline)
            max_y = ym + 1

    sx = first[0]
    sy = first[1]

    ex = second[0]
    ey = second[1]

    while sx != ex or sy != ey:
        puzzle[sy][sx] += 1
        sx += x_dir
        sy += y_dir
    puzzle[sy][sx] += 1

# pt2

print(calc_overlap(puzzle))
