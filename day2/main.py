# This is a sample Python script.

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.


def print_hi(name):
    # Use a breakpoint in the code line below to debug your script.
    print(f'Hi, {name}')  # Press ⌘F8 to toggle the breakpoint.


def validator(bounds, character, line):
    count = line.count(character)
    if int(bounds[0]) <= count <= int(bounds[1]):
        return True
    return False


def input_processor(password):
    print(password)
    var = password.replace(":", "")
    params = var.split()
    rules, character, line = params[0].split("-"), params[1], params[2]
    return validator(rules, character, line)


def read_file(filename, validator):
    file = open(filename, 'r')
    passwords = file.readlines()
    valid_passwords = [password for password in passwords if validator(password)]
    print(len(valid_passwords))


# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    file_name = "input"
    read_file(file_name, input_processor)
    print_hi('PyCharm')

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
