data = open("input.txt")

sum = 0

for game in data:
    game = game[game.index(":")+1::] # remove "Game x:" label
    rounds = [round.strip() for round in game.split(";")]
    numCubes = {
        "red": 0,
        "green": 0,
        "blue": 0
    }
    for round in rounds:
        colors = [color.strip() for color in round.split(",")]
        for color in colors:
            count, name = color.split()
            count = int(count)
            numCubes[name] = max(numCubes[name], count)
    sum += numCubes["red"] * numCubes["green"] * numCubes["blue"]

print(sum)
