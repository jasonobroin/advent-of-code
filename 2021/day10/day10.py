# https://adventofcode.com/2021/day/10
#
# Read the input from stdin

import sys

map_to_close = {'(': ')', '[': ']', '{':'}', '<': '>'}
map_to_points = {')': 3, ']': 57, '}': 1197, '>': 25137}

def valid_open(open):
    return open in ['(', '[', '{', '<']

def check_line(line, offset, close_char):
    while offset < len(line):
        test = line[offset]
#        print('cl', offset, test, close_char)
        if test == close_char:
            return offset
        elif valid_open(test):
            offset = check_line(line, offset + 1, map_to_close[test])
            if offset == 0:
                return 0
        elif test != close_char:
            global score
            bad_pts = map_to_points[test]
#            print(test, ' means bad_pts = ', bad_pts)
            score += bad_pts
            return 0
        offset += 1
    return len(line)

data = []

for line in sys.stdin:
    line = line.strip()
    t = []
    for j in line:
        t.append(j)
    data.append(t)

print(data)

score = 0

for line in data:
    if check_line(line, 1, map_to_close[line[0]]) == 0:
        print(f'{line} is corrupt')
    else:
        print(f'{line} is incomplete')

print(score)