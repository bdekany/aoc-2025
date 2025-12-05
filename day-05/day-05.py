#! /usr/bin/python3

fresh_interval = list()
ingredients = list()

def isInInterval(i, d):
    if i in d:
        return True
    return False

# Ouvre le fichier en mode lecture
nom_fichier = 'puzzle-input.txt'
with open(nom_fichier, 'r') as fichier:
    # Lit chaque ligne du fichier
    for ligne in fichier:
        if "-" in ligne:
            start, end = ligne.strip().split("-")
            fresh_interval.append(range(int(start), int(end)))
        
        elif len(ligne)>1:
            ingredients.append(int(ligne.strip()))

fresh_interval.sort(key=lambda r: r.start)
fusion_interval = list()

for r in fresh_interval:
    if len(fusion_interval) == 0:
        fusion_interval.append(r)
        continue
    
    if fusion_interval[-1].stop >= r.start:
        fusion_interval[-1] = range(fusion_interval[-1].start, max(fusion_interval[-1].stop, r.stop))
    else:
        fusion_interval.append(r)

result_part1 = 0
for ing in ingredients:
    for inter in fusion_interval:
        if isInInterval(ing, inter):
            result_part1 += 1
            break


result_part2 = 0
for inter in fusion_interval:
    result_part2 += len(inter) +1



print("Part1: ", result_part1)
print("Part2: ", result_part2)

