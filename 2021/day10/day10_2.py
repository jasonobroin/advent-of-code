# https://adventofcode.com/2021/day/10
#
# Read the input from stdin

import sys

map_to_close = {'(': ')', '[': ']', '{': '}', '<': '>'}
map_to_points = {')': 1, ']': 2, '}': 3, '>': 4}

def valid_open(open):
    return open in ['(', '[', '{', '<']

def check_line(line, offset, close_char, depth):
    global score
    while offset < len(line):
        test = line[offset]
#        print('cl', offset, test, close_char)
        if test == close_char:
#            print(test, offset)
            if depth != 0:
                return offset
            else:
                # A bit of a hack!
                depth -= 1
        elif valid_open(test):
            offset = check_line(line, offset + 1, map_to_close[test], depth + 1)
            # This is the path we're going through that's not accounting correctly
            if offset == 0:
                return 0
        elif test != close_char:
            # corrupt case
            return 0
        offset += 1
#    print(close_char)
    if depth != -1:
        score = score * 5 + map_to_points[close_char]
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
results = []

for line in data:
    score = 0
    if check_line(line, 1, map_to_close[line[0]], 0) == 0:
        print(f'{line} is corrupt')
    else:
        print(f'{line} is incomplete score {score}')
        results.append(score)

results.sort()
num = len(results) // 2
print(results[num])
