# input_file = open("example.txt") 
input_file = open("input.txt") 

lines = input_file.readlines()
input_file.close()

GRID_WIDTH = len(lines[0])
GRID_HEIGHT = len(lines)

# walk the grid backwards and forwards from the given position,
# then return the integer value of the enclosed string
def get_val(r: int, c: int) -> int:
    # print('getting val at', r, c)
    line = lines[r]

    left = c
    while left > 0 and line[left - 1].isdigit():
        left -= 1

    right = c + 1
    while right < len(line) - 1 and line[right].isdigit():
        right += 1

    # print(line[left:right])
    return int(line[left:right])

sum = 0

for i in range(GRID_HEIGHT):
    for j in range(GRID_WIDTH):
        char = lines[i][j]
        if char == '*':
            top_left_is_number = i >= 1 and j >= 1 and lines[i - 1][j - 1].isdigit()
            top_right_is_number = i >= 1 and j < len(lines[i]) and lines[i - 1][j + 1].isdigit()
            top_middle_is_number = i >= 1 and lines[i - 1][j - 1].isdigit()
            top_middle_is_dot = i >= 1 and lines[i - 1][j] == '.'
            bot_left_is_number = i < len(lines) - 1 and j >= 1 and lines[i + 1][j - 1].isdigit()
            bot_right_is_number = i < len(lines) - 1 and j < len(lines[i]) and lines[i + 1][j + 1].isdigit()
            bot_middle_is_number = i < len(lines) - 1 and lines[i + 1][j - 1].isdigit()
            bot_middle_is_dot = i < len(lines) - 1 and lines[i + 1][j] == '.'
            # print('found a *')
            # print(top_left_is_number, top_right_is_number, top_middle_is_dot)

            nearby_nums = []

            # check left
            if j >= 1 and lines[i][j - 1].isdigit():
                nearby_nums.append(get_val(i, j - 1))
            # check right
            if j < len(lines[i]) - 1 and lines[i][j + 1].isdigit():
                nearby_nums.append(get_val(i, j + 1))
            # check row above
            if top_left_is_number and top_right_is_number and top_middle_is_dot:
                nearby_nums.append(get_val(i - 1, j - 1))
                nearby_nums.append(get_val(i - 1, j + 1))
            elif top_left_is_number:
                nearby_nums.append(get_val(i - 1, j - 1))
            elif top_middle_is_number:
                nearby_nums.append(get_val(i - 1, j))
            elif top_right_is_number:
                nearby_nums.append(get_val(i - 1, j + 1))
            # check row below
            if bot_left_is_number and bot_right_is_number and bot_middle_is_dot:
                nearby_nums.append(get_val(i + 1, j - 1))
                nearby_nums.append(get_val(i + 1, j + 1))
            elif bot_left_is_number:
                nearby_nums.append(get_val(i + 1, j - 1))
            elif bot_middle_is_number:
                nearby_nums.append(get_val(i + 1, j))
            elif bot_right_is_number:
                nearby_nums.append(get_val(i + 1, j + 1))

            if len(nearby_nums) == 2:
                sum += nearby_nums[0] * nearby_nums[1]

print("FINAL SUM:")
print(sum)
