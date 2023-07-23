with open("input.txt") as f:
    content = f.readlines()
    

listofsums = []
numsum = 0
for i in content:
    if i == "\n":
        listofsums.append(numsum)
        numsum = 0
        continue
    numsum = numsum + int(i)

maxcal = max(listofsums)
print(maxcal)
listofsums.sort(reverse=True)

backup = sum(listofsums[:3])
print(backup)

