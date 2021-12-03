# https://adventofcode.com/2021/day/3
#
# Read the input from stdin

import sys


def count_bits(array):
    count = [0] * digits

    # Count number of bits that are set
    for num in array:
        for bit in range(digits - 1, -1, -1):
            if num & 1 << bit:
                count[bit] += 1

    return count


def reduce(array, test):
    # We have our bit counts, so run through this
    bitpos = digits - 1
    while len(array) > 1:
        count = count_bits(array)
        val = count[bitpos]
        half = len(array) / 2
        keep = test
        if val >= half:
            keep = keep ^ 1
        array = [i for i in array if ((i >> bitpos) & 1 == keep)]
        # print(val, half, keep, array)
        bitpos -= 1
    return array

puzzle = []

digits = 0
for line in sys.stdin:
    digits = len(line)
    puzzle.append(int(line, 2))

# pt1

lines = len(puzzle)
count = count_bits(puzzle)

# Note our array starts with bit 0 - parse it to build up our 'gamma'
# We can see if there are more bits than half the number of lines to indicate if '1' is the common bit

gamma = 0
idx = 0
for num in count:
    if num > lines/2:
        gamma |= 1 << idx
    idx += 1

# epsilon is based on the least common bit - we can NOT & mask
epsilon = ~gamma & ((1 << digits) - 1)

print(count)
print(gamma)
print(epsilon)
print(f'Power Consumption =  {gamma * epsilon}')

# pt2

o2 = reduce(puzzle, 0)
co2 = reduce(puzzle, 1)

print(f'Oxygen generator rating = {o2[0] * co2[0]}')
