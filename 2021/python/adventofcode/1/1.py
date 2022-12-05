input = open('input.txt', 'r')
lines = input.readlines()

increments = 0
maxvalue = int(lines[0])
for line in lines:
    element  = int(line)
    if element > maxvalue:
        increments += 1
    maxvalue = element
print(increments, end='')
