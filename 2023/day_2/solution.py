# AoC 2023 - Day 2

def parse_line(line: str) -> dict:
    line = line.strip()
    ret, game = {}, {}
    game_id, rest = line.split(':')
    ret['id'] = int(game_id.split(' ')[1])
    for i, l in enumerate(rest.split(';')):
        game[f'{i}'] = l
    ret['games']= game
    return ret

def check_bag(games: dict) -> bool:
    for key,val in games.items():
        splitted = val.lstrip().split(',')
        for value in splitted:
            value = value.strip().split()
            if value[1] == 'red' and int(value[0]) > 12:
                return False
            elif value[1] == 'green' and int(value[0]) > 13:
                return False
            elif value[1] == 'blue' and int(value[0]) > 14:
                return False
    return True

def fewest_number_of_cubes(games: dict) -> int:
    red, green, blue = 0, 0, 0
    for key,val in games.items():
        splitted = val.lstrip().split(',')
        for value in splitted:
            value = value.strip().split()
            if value[1] == 'red':
                red = max(red, int(value[0]))
            elif value[1] == 'green':
                green = max(green, int(value[0]))
            elif value[1] == 'blue':
                blue = max(blue, int(value[0]))
    return red * green * blue

def part_1(fname: str) -> int:
    s = 0
    with open(fname) as f:
        for line in f:
            x = parse_line(line)
            i_d = x['id']
            if check_bag(x['games']):
                s += i_d
    return s

def part_2(fname: str) -> int:
    s = 0
    with open(fname) as f:
        for line in f:
            x = parse_line(line)
            s += fewest_number_of_cubes(x.get('games'))
    return s

if __name__ == "__main__":
    filename = 'input_1.txt'
    print(part_1(filename))
    print(part_2(filename))