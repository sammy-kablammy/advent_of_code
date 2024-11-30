DIGITS = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']

running_total = 0
input_file = open("input.txt")
for line in input_file:
    left_digit = ''
    for char in line:
        if char in DIGITS:
            left_digit = char
            break
    right_digit = ''
    for char in line[::-1]:
        if char in DIGITS:
            right_digit = char
            break
    running_total += int(left_digit + right_digit)
input_file.close()

print(f"The sum of the calibration values is {running_total}")
