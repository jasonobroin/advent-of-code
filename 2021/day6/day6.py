# https://adventofcode.com/2021/day/6
#
# Read the input from stdin

import sys

line = sys.stdin.readline()
data = [int(s) for s in line.strip().split(sep=',')]
print(data)

max_ages = 9
ages = [ 0 ] * max_ages
for i in data:
    ages[i] += 1

print(ages)

# pt1

count = 0
while (count < 80):
    num = len(data)
    for i in range(0, num):
        data[i] -= 1
        if data[i] == -1:
            data[i] = 6
            data.append(8)
    count += 1

print(count, ":", len(data))

# pt2

# We can't brut force - it takes forever and our array will eventually exceed all memory
# We can take a different approach - there are only 8 lifetimes we care about; we just need to know
# how many lattern fish are of the same age, and decrement them together. When a set reaches -1, we
# add that many more to the new born group

print("")
print(ages)
count = 0
while (count < 256):
    newborn = 0
    for i in range(0, max_ages):
        if i == 0:
            newborn = ages[0]
        else:
            ages[i - 1] = ages[i]
    ages[max_ages-1] = newborn
    ages[6] += newborn
    count += 1
    if count == 80 or count == 256:
        print(count, ": ", sum(ages))

