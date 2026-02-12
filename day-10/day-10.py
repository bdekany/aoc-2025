#! /usr/bin/python3

class Machine:
    def __init__(self, raw):
        self.joltage = raw.pop()

        self.switch = list()
        while len(raw) > 1:
            tmp = self.__clean(raw.pop())
            self.switch.append( self.__to_array(tmp) )

        self.lights = list()
        for i in self.__clean(raw.pop()):
            if i == "#":
                self.lights.append(1)
            else:
                self.lights.append(0)

    def __clean(self, a):
        return a[1:len(a)-1]
    
    def __to_array(self, a):
        return list(map(lambda x: int(x), a.split(",")))

    def print(self):
        print("light:", self.lights, "switch", self.switch)


all_machines = list()

# Ouvre le fichier en mode lecture
nom_fichier = 'puzzle-input.txt'
with open(nom_fichier, 'r') as fichier:
    # Lit chaque ligne du fichier
    for ligne in fichier:
        tmp = ligne.strip().split(" ")
        all_machines.append(Machine(tmp))

for m in all_machines:
    m.print()
    temoin = [0] * len(m.lights)
    print(temoin)
    for s in m.switch:
        for b in s:
            temoin[b] = 1 - temoin[b]
        if temoin == m.lights:
            print(temoin)
            break