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


# pt2

rel_offsets = { (-1,0), (+1, 0), (0, -1), (0, +1), }

total_vals = len(data)

basins = []

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
            # We've found a low point - now we need to find the number of values around it bound by 9
            chk_mtx = [row[:] for row in data] # copy array

            def check_around(yo, xo):
                count = 0
                for off in rel_offsets:
                    rx = xo + off[0]
                    ry = yo + off[1]
                    if rx >= 0 and rx < x_max and ry >= 0 and ry < y_max:
                        if chk_mtx[ry][rx] != 9:
                            count += 1
                            chk_mtx[ry][rx] = 9
                            count += check_around(ry, rx)
                return count


            count = 0
            count += check_around(y_offset, x_offset)

            # print(y_offset, x_offset, count)
            basins.append(count)

basins.sort(reverse=True)
# print(basins)
print("prod", basins[0] * basins[1] * basins[2])