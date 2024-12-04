# today might be too much for my awking ability. resorting to python...

import sys

text = sys.stdin.read()
output = ""
enabled = True

# this is the most hacky piece of garbage of all time
for i in range(len(text)):
    if text[i:i+len("do()")] == "do()":
        enabled = True
    if text[i:i+len("don't()")] == "don't()":
        enabled = False
    if enabled:
        output += text[i]

print(output)
