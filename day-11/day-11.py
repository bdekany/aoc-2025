#! /usr/bin/python3


def go_deep(devices, node, path):
    local = 0
    for out in devices[node]:
        print("test:", out)
        if out == "out":
            print(path)
            return 1
        else:
            path.append(out)
            local += go_deep(devices, out, path)
        path.pop()
    return local

def go_up(devices, node, path):
    local = 0
    for d in devices:
        print("test:", d)
        if d == "svr":
            return 1
        if d in path:
            return 0
        if node in devices[d]:
            path.append(d)
            print(path)
            local += go_up(devices, d, path)
        # path.pop()
    return local

devices = {}

# Ouvre le fichier en mode lecture
nom_fichier = 'puzzle-input.txt'
with open(nom_fichier, 'r') as fichier:
    # Lit chaque ligne du fichier
    for ligne in fichier:
        tmp = ligne.strip().split(":")
        devices[tmp[0]] = tmp[1].strip().split(" ")
        
resultat_part1 = 0

current = "you"
path = list()
path.append(current)
# resultat_part1 = go_deep(devices, current, path)

current = "dac"
path = list()
path.append(current)
resultat_part1 = go_up(devices, current, path)

print("resultat", resultat_part1)