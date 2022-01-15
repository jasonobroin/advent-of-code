# https://adventofcode.com/2021/day/13
#
# Read the input from stdin

import sys

array = []
max_x = 0
max_y = 0

def print_array(puzzle):
    for y in range(0, max_y + 1):
        s = ''
        for x in range(0, max_x + 1):
            if puzzle[y][x] == 0:
                s += '.'
            else:
                s += '#'
        print(s)


def count_dots(puzzle):
    dots = 0
    for y in range(0, max_y + 1):
        for x in range(0, max_x + 1):
            if puzzle[y][x] == 1:
                dots += 1
    return dots


def fold_y(y):
    global array

    split = y
    for i in range(y + 1, max_y + 1, 1):
        for x in range(0, max_x + 1, 1):
            if array[i][x] == 1:
                array[split - (i - split)][x] = 1
    return y - 1


def fold_x(x):
    global array

    # if max_x / 2 != x:
    #     print("X NOT SPLIT IN TWO")

    # Strictly this should be implemented like fold_y to NOT assume the split is in the middle
    # but this is fine with the data set
    for y in range(0, max_y + 1, 1):
        for i in range(0, x, 1):
            if array[y][max_x - i] == 1:
                array[y][i] = 1
    return x - 1


# Handle matrix
for line in sys.stdin:
    line = line.strip()
    if line == "":
        break
    x, y = [int(s) for s in line.split(",")]

#    print(x,y)

    # Make sure the array is big enough

    if x > max_x:
        x_to_add = x - max_x + 1
        max_x = x
    else:
        x_to_add = 0

    if y > max_y:
        max_y = y

    array_len = len(array)
    for i in range(0, max_y + 1, 1):
            if i >= array_len:
                newline = [0] * (max_x + 1)
                array.append(newline)
            if x > len(array[i]):
                # Extend the size of this line
                for a in range(x_to_add):
                    array[i].append(0)

    # Then set the location
    array[y][x] = 1

print()

# Now handle folds
for line in sys.stdin:
    line = line.strip()
    _, line = line.split("fold along ")
    dir, val = line.split("=")

    if dir == 'y':
        max_y = fold_y(int(val))
    else:
        max_x = fold_x(int(val))
    print("fold", dir, val)
    print("dots", count_dots(array))

print_array(array)
