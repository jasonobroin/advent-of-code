# https://adventofcode.com/2021/day/1
#
# Read the input from stdin

def count_incs(array):
    larger = 0
    prev = 99999
    for val in array:
        if val > prev:
            larger += 1
        prev = val
    return larger


import sys

depths = []

for line in sys.stdin:
    depths.append(int(line))

# pt1

print(count_incs(depths))

# pt2

three_depths = []
offset = 0

while offset < len(depths) - 2:
    total = depths[offset] + depths[offset+1] + depths[offset+2]
    three_depths.append(total)
    offset += 1

print(count_incs(three_depths))




