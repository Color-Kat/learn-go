part1 = list(map(int, input().split()))
n = part1[0]
whitePrice = part1[1]
blackPrice = part1[2]
costumes = list(map(int, input().split()))

l, r = 0, n - 1

# 5 100 1
# 0 1 2 1 2

price = 0

while l <= r:
    if costumes[l] == costumes[r] or (costumes[l] == 2 or costumes[r] == 2):
        if costumes[l] == 2 and costumes[r] == 2:
            count = 2
            if l == r:
                count = 1
            price += min(whitePrice, blackPrice) * count
        else:
            # l - is free
            if costumes[l] == 2:
                if costumes[r] == 1:
                    price += blackPrice
                if costumes[r] == 0:
                    price += whitePrice

            # r - is free
            if costumes[r] == 2:
                if costumes[l] == 1:
                    price += blackPrice
                if costumes[l] == 0:
                    price += whitePrice

        r -= 1
        l += 1
    else:
        price = -1
        break

print(price)