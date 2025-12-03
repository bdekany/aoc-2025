#! /usr/bin/python3

combinaisons = list()

# Ouvre le fichier en mode lecture
nom_fichier = 'puzzle-input.txt'
with open(nom_fichier, 'r') as fichier:
    # Lit chaque ligne du fichier
    for ligne in fichier:
        # Creation carte
        combinaisons.append(ligne.strip())

# Position Initial
position = 50

# Counters
puzzle_part1 = 0
puzzle_part2 = 0

for i in range(len(combinaisons)):
    # Slit de chaque ligne
    operator = combinaisons[i][0]
    distance = int( combinaisons[i][1:] )

    if operator == "R":
        puzzle_part2 += ( position + distance ) // 100
        position = ( position + distance ) % 100
    else:
        if position == 0:
            puzzle_part2 -= 1

        puzzle_part2 += abs ( ( position - distance ) // 100 )
        position = ( position - distance ) % 100
        
        if position == 0:
            puzzle_part2 += 1

    if position == 0:
        puzzle_part1 += 1

print("Result part 1:", puzzle_part1)
print("Result part 2:", puzzle_part2)
