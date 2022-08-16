s1 = input()
s2 = input()

m = {}

for c in s2:
    if c in m.keys():
        m[c] += 1
    else:
        m[c] = 1

queue = []
        
def is_over():
    for v in m.values():
        if v > 0:
            return False
    return True
    
for i in range(len(s1)):
    c = s1[i]
    if c in m.keys():
        queue.append(i)
        m[c] -= 1
        if is_over():
            break
        while m[c] < 0 and s1[queue[0]] == c:
            queue = queue[1:]
            m[c] += 1
    
print(s1[queue[0]: queue[-1]+1])