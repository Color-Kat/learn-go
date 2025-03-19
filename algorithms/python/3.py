part1 = list(map(int, input().split()))
n = part1[0]
m = part1[1]
lst = list(map(int, input().split()))
result = []

for j in range(m):
    fromI = int(input())-1
    counts = {}

    for i in range(fromI, n):
        counts[lst[i]] = counts.get(lst[i], 0) + 1

    print(len(counts))
