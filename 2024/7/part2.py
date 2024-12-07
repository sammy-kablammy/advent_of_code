# resorting to python again; itertools.product is too convenient

import sys
import itertools


def solve(line: str) -> int:
    tokens = line.split()
    target = int(tokens[0][:-1])
    values = [int(v) for v in tokens[2::]]

    prod = itertools.product(["+", "*", "|"], repeat=len(values))
    for p in prod:
        total = int(tokens[1])
        for i in range(len(values)):
            if p[i] == "+":
                total += values[i]
            elif p[i] == "*":
                total *= values[i]
            elif p[i] == "|":
                total = int(str(total) + str(values[i]))
            else:
                print("unexpected operator", file=sys.stderr)
        if total == target:
            return total

    return 0


sum = 0
for line in sys.stdin.readlines():
    sum += solve(line)
print(sum)
