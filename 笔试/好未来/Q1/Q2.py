def func(s):
    map = {}
    for c in s:
        if c in map:
            map[c] += 1
        else:
            map[c] = 1
    for k in map.keys():
        print(k, map[k])

def main():
    func(input())

if __name__ == '__main__':
    main()