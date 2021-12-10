# https://adventofcode.com/2021/day/9
#
# Read the input from stdin

import sys

data = []

for line in sys.stdin:
    line = line.strip()
    t = []
    for j in line:
        t.append(int(j))
    data.append(t)

x_max = len(data[0])
y_max = len(data)

# print(x_max, y_max, data)

# pt1

rel_offsets = { (-1,0), (+1, 0), (0, -1), (0, +1) }
#rel_offsets = { -1, +1, -x_max -1, -x_max, -x_max + 1, x_max - 1, x_max, x_max + 1 }

total_vals = len(data)

sum = 0

x_offset = 0
y_offset = 0
for y_offset in range(0, y_max):
    for x_offset in range(0, x_max):
        # Check numbers around us
        lowest = True
        for off in rel_offsets:
            rx = x_offset + off[0]
            ry = y_offset + off[1]
            if rx >= 0 and rx < x_max and ry >= 0 and ry < y_max:
                # The offset is in the array
                if data[ry][rx] <= data[y_offset][x_offset]:
                    lowest = False

        if lowest:
            # print("lowest at offset", offset, data[offset])
            sum += 1 + data[y_offset][x_offset]

print(sum)
