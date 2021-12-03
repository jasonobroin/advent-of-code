# https://adventofcode.com/2021/day/2
#
# Read the input from stdin

import sys

puzzle = []

for line in sys.stdin:
    puzzle.append(str.split(line))

# pt1

hor = 0
depth = 0

for l in puzzle:
    if l[0] == 'forward':
        hor += int(l[1])
    elif l[0] == 'up':
        depth -= int(l[1])
    elif l[0] == 'down':
        depth += int(l[1])
    else:
        print("bad input")

print(f"depth = {depth} hor = {hor}")

print("pt1 = ", depth * hor)

# pt2

hor = 0
depth = 0
aim = 0

for l in puzzle:
    if l[0] == 'forward':
        hor += int(l[1])
        depth += aim * int(l[1])
    elif l[0] == 'up':
        aim -= int(l[1])
    elif l[0] == 'down':
        aim += int(l[1])
    else:
        print("bad input")

print(f"depth = {depth} hor = {hor} aim = {aim}")

print("pt2 = ", depth * hor)

