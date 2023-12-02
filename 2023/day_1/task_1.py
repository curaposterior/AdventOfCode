# AoC 2023 - Day 1

def task_1(fname: str):
    sum_a = 0
    with open(fname) as f:
        for i in f:
            i = i.strip()
            first, last = None, None
            for char in i:
                if char.isnumeric() and first is None:
                    first = char
                elif char.isnumeric():
                    last = char
            if first and last:
                sum_a += int(first + last)
            else:
                sum_a += int(first + first)
    print(sum_a)


def parse_first(line):
    keywords = {'one': '1', 'two': '2', 'three': '3', 'four': '4', 'five': '5',
                 'six':'6', 'seven': '7', 'eight': '8', 'nine': '9'}
    for i, v in enumerate(line):
        if v.isdigit():
            return v

        for k in keywords:
            if line[i:i + len(k)] == k:
                return keywords[k]

def parse_last(line):
    keywords = {'one': '1', 'two': '2', 'three': '3', 'four': '4', 'five': '5',
                 'six':'6', 'seven': '7', 'eight': '8', 'nine': '9'}
    for i in range(len(line)-1, -1, -1):
        if line[i].isdigit():
            return line[i]
        for k in keywords:
            if line[i-len(k) + 1:i + 1] == k:
                return keywords[k]


def task_2(fname: str):
    global_sum = 0
    with open(fname) as f:
        for line in f:
            line = line.strip()

            first = parse_first(line)
            last = parse_last(line)
            global_sum += int(first + last)
    print(global_sum)
                    

if __name__ == '__main__':
    task_1('input_1.txt')
    task_2('input_1.txt')
