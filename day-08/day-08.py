#! /usr/bin/python3

from math import sqrt

def v_distance(a, b):
    x = pow(b[0] - a[0], 2)
    y = pow(b[1] - a[1], 2)
    z = pow(b[2] - a[2], 2)
    return sqrt(x + y + z)

junction_boxes = list()

# Ouvre le fichier en mode lecture
nom_fichier = 'puzzle-input.txt'
with open(nom_fichier, 'r') as fichier:
    # Lit chaque ligne du fichier
    for ligne in fichier:
        tmp = ligne.strip().split(",")
        junction_boxes.append( list(map(lambda x: int(x), tmp)) )

print("List Junctions Boxes")
print(junction_boxes)

vector_diff = dict()

for j in range(len(junction_boxes)):
    for c in range(j+1, len(junction_boxes)):
        a = junction_boxes[j]
        b = junction_boxes[c]
        vector_diff[str(a),str(b)] = v_distance(a, b)

print("Sorted by values:")
# Sort the dictionary by values and print the key-value pairs
sorted_items = sorted(vector_diff.items(), key=lambda kv: kv[1])
print(sorted_items[:5])

super_set = list()
nmb_cable = 9

for i in sorted_items[:10]:
    k, v = i[0]
    print("---")
    print("New set:", k, v)

    if len(super_set) == 0:
        print("Init avec ", k, v)
        super_set.append(set([v, k]))
        nmb_cable -= 1
        continue

    a = b = None

    for s in super_set:
        if k in s:
            print("k found")
            a = s
        
        if v in s:
            print("v found")
            b = s

    if a is None and b is None:
        print("found none")
        super_set.append(set([v, k]))
        nmb_cable -= 1
        continue

    if a == b:
        print("Circuit Connu")
        continue

    if a and b:
        tmp = a.union(b)
        super_set.remove(a)
        super_set.remove(b)
        super_set.append(tmp)
        print("fusion:", tmp)
        nmb_cable -= 1
    elif a and not b:
        a.add(v)
        print("ajout v", a)
        nmb_cable -= 1
    elif b and not a:
        b.add(k)
        print("ajout k", b)
        nmb_cable -= 1
    
    if nmb_cable == 0:
        print("fin")
        break
    


print("circuits")
print(super_set)

size = list()

for ss in super_set:
    size.append(len(ss))

result = sorted(size)
print(result)
print("Result Part1", result[-3]*result[-2]*result[-1])