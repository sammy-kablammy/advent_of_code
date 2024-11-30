import sys

input_file = open(sys.argv[1])
lines = input_file.readlines()
input_file.close()

total_score = 0

for line in lines:
    idx1 = line.index(':')
    idx2 = line.index('|')
    winning_numbers = line[idx1+1:idx2].strip().split()
    my_numbers = line[idx2+1:].strip().split()
    # print(winning_numbers)
    # print(my_numbers)
    count = 0
    for my_num in my_numbers:
        if my_num in winning_numbers:
            count += 1
    if count > 0:
        total_score += 2 ** (count - 1)

print('total score is', total_score)
