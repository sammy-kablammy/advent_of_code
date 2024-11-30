# input_file = open("example.txt") 
input_file = open("input.txt") 

# bruh the numbers at the end of lines got me big time

rawlines = input_file.readlines()
input_file.close()

# literally why did i do this in the beginning
# lines = [s[0:-1:1] for s in rawlines] # remove \n's

lines = rawlines

GRID_WIDTH = len(lines[0])
GRID_HEIGHT = len(lines)

# instead of finding numbers, then checking if those numbers are valid...
# find the SYMBOLS, then mark adjacent squares as valid

# first pass: find all the symbols and mark their adjacent squares as valid
valid = [[False for _ in l] for l in lines]
symbols = ['@', '#', '$', '%', '&', '*', '-', '=', '/', '+']
# symbols = '@#$%&*-=+/'
def markvalid(r: int, c: int) -> None:
    for roffset in (-1, 0, 1):
        for coffset in (-1, 0, 1):
            rAdj = r + roffset
            cAdj = c + coffset
            if rAdj < GRID_HEIGHT and cAdj < GRID_WIDTH:
                valid[rAdj][cAdj] = True

for i in range(GRID_HEIGHT):
    for j in range(GRID_WIDTH):
        char = lines[i][j]
        if char in symbols:
            markvalid(i, j)

sum = 0

# second pass: count up them numbers!
for i in range(GRID_HEIGHT):
    num_so_far = ''
    is_current_number_valid = False
    for j in range(GRID_WIDTH):
        char = lines[i][j]
        if char.isdigit() and num_so_far == '': # start of number
            num_so_far += char
            is_current_number_valid = is_current_number_valid or valid[i][j]
        elif char.isdigit() and num_so_far != '': # middle of number
            num_so_far += char
            is_current_number_valid = is_current_number_valid or valid[i][j]
        elif not(char.isdigit()) and num_so_far != '': # end of number
            if is_current_number_valid:
                sum += int(num_so_far)
                print('so far:' + num_so_far)
            else:
                print('elsers.', num_so_far)
            num_so_far = ''
            is_current_number_valid = False
        else: # no number
            pass

print("FINAL SUM:")
print(sum)
