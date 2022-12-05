def read_file():
    input_file = open('input.txt', 'r')
    lines = input_file.readlines()
    return lines

def apply_instruction(x_position,y_position,movement,units):
    if movement == 'forward':
        x_position += units
    if movement == 'down':
        y_position += units
    if movement == 'up':
        y_position -= units
    return x_position, y_position

def apply_instruction_with_aim(x_position,y_position,movement,units,aim):
    if movement == 'forward':
        y_position += (aim * units)
        x_position += units
    if movement == 'down':
        aim += units
    if movement == 'up':
        aim -= units
    return x_position, y_position, aim

def main():
    x_position = 0
    y_position = 0
    aim = 0
    instructions = read_file()
    for instruction in instructions:
        instruction = instruction.split()
        movement, units = instruction[0], int(instruction[1])
        x_position, y_position, aim = apply_instruction_with_aim(
                x_position, 
                y_position, 
                movement, 
                units, 
                aim)
    print(x_position, y_position, x_position * y_position)

main()
