part1 = list(map(int, input().split()))
n = part1[0]
t = part1[1]
c = part1[2]
severity = list(map(int, input().split()))

waysCount = 0

pref = [0] * (n+1)

for i in range(len(severity)):
    pref[i + 1] = pref[i] + severity[i]

    if i < c - 1:
        continue

    print(pref[i+1] - pref[i-c+1], "=", ((pref[i+1] - pref[i-c+1]) / c))
    if (pref[i+1] - pref[i-c+1]) <= t*c:
        waysCount += 1


    # if pref[i+c] - pref[i] <= t:
    #     waysCount += 1

print(pref)


print(waysCount)