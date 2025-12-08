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
nmb_cable = 11

for i in sorted_items:
    found = False
    k, v = i[0]
    print("New set:", k, v)

    if len(super_set) == 0:
        print("Init avec ", k, v)
        super_set.append(set([v, k]))
        nmb_cable -= 1
        continue

    for s in super_set:
        if k in s and v in s:
            print("Circuit connu", k, v)
            found = True
            continue

        elif k in s:
            print("found k")
            s.add(v)
            print(s)
            nmb_cable -= 1
            found = True
            tmp = s
            continue
        
        elif v in s:
            print("found v")
            if tmp:
                tmp = tmp.union(s)
                print(tmp)
                del(s)
            else:
                s.add(k)
                print(s)
            nmb_cable -= 1
            found = True
            continue

    tmp = None

    if not found:
        print("found none")
        super_set.append(set([v, k]))
        nmb_cable -= 1
    
    if nmb_cable == 0:
        break

print("circuits")
print(super_set)

size = list()

for ss in super_set:
    size.append(len(ss))

result = sorted(size)
print("Result Part1", result[-3]*result[-2]*result[-1])