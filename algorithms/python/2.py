part1 = list(map(int, input().split()))
n = part1[0]
t = part1[1]
c = part1[2]
severity = list(map(int, input().split()))

# Use sliding windows
l, r = 0, c - 1

waysCount = 0

for i in range(len(severity)-c + 1):
    verdict = True
    for j in range(c):
        if severity[i+j] > t:
            verdict = False
            break

    if verdict:
        waysCount += 1

print(waysCount)