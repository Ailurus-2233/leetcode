def func(s: str):
    '''
    最长子串，其中字串中的字符出现次数为偶数次
    '''
    n = len(s)
    pos = [-1 for x in range(1 << 26)]
    ans = 0
    status = 0
    pos[0] = 0
    for i in range(n):
        ch = ord(s[i]) - ord('a')
        status ^= 1 << ch
        if pos[status] != -1:
            ans = max(ans, i + 1 - pos[status])
        else:
            pos[status] = i + 1
    return ans


if __name__ == '__main__':
    s = input()
    print(func(s))