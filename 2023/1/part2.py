DIGITS = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
WORDS = ['one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']

running_total = 0
input_file = open("input.txt")
for line in input_file:
    digits_found_on_this_line = []
    for i in range(len(line)):
        if line[i] in DIGITS:
            digits_found_on_this_line.append(line[i])
        elif line[i:i+3] in WORDS:
            digits_found_on_this_line.append(str(WORDS.index(line[i:i+3]) + 1))
        elif line[i:i+4] in WORDS:
            digits_found_on_this_line.append(str(WORDS.index(line[i:i+4]) + 1))
        elif line[i:i+5] in WORDS:
            digits_found_on_this_line.append(str(WORDS.index(line[i:i+5]) + 1))
    # print(int(digits_found_on_this_line[0] + digits_found_on_this_line[-1]))
    running_total += int(digits_found_on_this_line[0] + digits_found_on_this_line[-1])
input_file.close()

print(f"The sum of the calibration values is {running_total}")
