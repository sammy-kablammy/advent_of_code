# i know this is extremely slow lol

import sys

input_file = open(sys.argv[1])
lines = input_file.readlines()
input_file.close()

total_cards = len(lines)

copies = {}
for i in range(len(lines)):
    copies[i + 1] = 1

for line in lines:
    card_number = int(line[5:line.index(':')])
    for copy_num in range(copies[card_number]):
        idx1 = line.index(':')
        idx2 = line.index('|')
        winning_numbers = line[idx1+1:idx2].strip().split()
        my_numbers = line[idx2+1:].strip().split()
        count = 0
        for my_num in my_numbers:
            if my_num in winning_numbers:
                count += 1
        for i in range(count):
            copies[card_number + i + 1] += 1
            total_cards += 1

print(copies)
print('total number of cards is', total_cards)
