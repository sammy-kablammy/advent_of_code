data = open("input.txt")

color_amounts = {
        "red": 12,
        "green": 13,
        "blue": 14
}

gamePossibilities = [False]

gameNumber = 0
for game in data:
    gameNumber += 1
    gamePossibilities.append(True)
    game = game[game.index(":")+1::] # remove "Game x:" label
    rounds = [round.strip() for round in game.split(";")]
    for round in rounds:
        colors = [color.strip() for color in round.split(",")]
        for color in colors:
            count, name = color.split()
            if int(count) > color_amounts[name]:
                gamePossibilities[gameNumber] = False

sum = 0
for i in range(len(gamePossibilities)):
    if gamePossibilities[i]:
        sum += i

print(sum)
