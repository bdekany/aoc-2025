#! /usr/bin/python3

def rect_area(a, b):
    x = abs(a[0] - b[0]) + 1
    y = abs(a[1] - b[1]) + 1
    return x * y

tiles_map = list()

# Ouvre le fichier en mode lecture
nom_fichier = 'puzzle-input.txt'
with open(nom_fichier, 'r') as fichier:
    # Lit chaque ligne du fichier
    for ligne in fichier:
        tmp = ligne.strip().split(",")
        tiles_map.append( list(map(lambda x: int(x), tmp)) )

print("List Tiles")
print(tiles_map)

vector_diff = dict()

for j in range(len(tiles_map)):
    for c in range(j+1, len(tiles_map)):
        a = tiles_map[j]
        b = tiles_map[c]
        vector_diff[str(a),str(b)] = rect_area(a, b)

print("Sorted by values:")
# Sort the dictionary by values and print the key-value pairs
sorted_items = sorted(vector_diff.items(), key=lambda kv: kv[1])
print("Resultat Part 1:", sorted_items[-1:][0][1])
