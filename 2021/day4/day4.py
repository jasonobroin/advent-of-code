# https://adventofcode.com/2021/day/4
#
# Read the input from stdin

import sys

puzzle = []

first_line = sys.stdin.readline()
numbers = [int(i) for i in first_line.split(sep=",")]

boards = []

w, h = 5, 5
entries = w * h

def check_winner(board):
    # check rows
    for idx in range(0, entries, w):
        val_total = 0
        state_total = 0
        for idx2 in range(idx, idx + w):
            val_total += board[idx2][0]
            state_total += board[idx2][1]
        if state_total == w:
            # Count unmarked numbers
            unmarked_total = 0
            for idx3 in range(0, entries):
                if board[idx3][1] == 0:
                    unmarked_total += board[idx3][0]
            return unmarked_total

    # check columns
    for idx in range(0, w, 1):
        val_total = 0
        state_total = 0
        for idx2 in range(idx, idx + w * (h-1) + 1, w):
            val_total += board[idx2][0]
            state_total += board[idx2][1]
        if state_total == h:
            # Count unmarked numbers
            unmarked_total = 0
            for idx3 in range(0, entries):
                if board[idx3][1] == 0:
                    unmarked_total += board[idx3][0]
            return unmarked_total

    return 0

# Main
while True:
    skip = sys.stdin.readline()
    if skip == '':
        break
    array = []
    for y in range(h):
        first_line = sys.stdin.readline()
        # Probably a better way to do this
        l = [[int(i), 0] for i in first_line.split()]
        for val in l:
            array.append(val)

    boards.append(array)

boards_complete = [False] * (len(boards)+1)

# Pick bingo numbers and score off in all hands
winner = 1
for num in numbers:
    board_num = 0
    for board in boards:
        if boards_complete[board_num]:
            board_num += 1
            continue
        try:
            idx = board.index([num, 0])
        except ValueError:
            board_num += 1
            continue
        if idx != entries:
            board[idx][1] = 1
        win = check_winner(board)
        if win != 0:
            print(f"winner {winner} = {win * num}")
            boards_complete[board_num] = True
            winner += 1
        board_num += 1


