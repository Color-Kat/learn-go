# Чтение данных
n = int(input())  # Размер массива
arr = list(map(int, input().split()))  # Сам массив

# Построение префиксных сумм
prefix_sum = [0] * (n + 1)  # Массив префиксных сумм, начинаем с 0
for i in range(n):
    prefix_sum[i + 1] = prefix_sum[i] + arr[i]  # Считаем префиксную сумму

# Пример: обработка запросов на сумму на отрезке [l, r]
q = int(input())  # Количество запросов
for _ in range(q):
    l, r = map(int, input().split())  # Индексы отрезка (1-based в условии)
    l -= 1  # Переводим в 0-based
    r -= 1
    # Сумма на отрезке [l, r] = prefix_sum[r + 1] - prefix_sum[l]
    result = prefix_sum[r + 1] - prefix_sum[l]
    print(result)

# ----------------

# Чтение данных
n, k = map(int, input().split())  # n - размер массива, k - целевая сумма
arr = list(map(int, input().split()))  # Сам массив

# Построение префиксных сумм
prefix_sum = [0] * (n + 1)
for i in range(n):
    prefix_sum[i + 1] = prefix_sum[i] + arr[i]

# Подсчёт количества подотрезков с суммой k
count = 0
sum_count = {0: 1}  # Словарь для хранения частоты префиксных сумм

for i in range(1, n + 1):
    # Если prefix_sum[i] - k есть в словаре, значит есть подотрезок с суммой k
    count += sum_count.get(prefix_sum[i] - k, 0)
    # Добавляем текущую префиксную сумму в словарь
    sum_count[prefix_sum[i]] = sum_count.get(prefix_sum[i], 0) + 1

print(count)