def binary_search(arr, target):
    left, right = 0, len(arr) - 1

    while left <= right:
        mid = (left + right) // 2
        if arr[mid] == target:
            return mid  # Нашли элемент, возвращаем его индекс
        elif arr[mid] < target:
            left = mid + 1  # Ищем в правой половине
        else:
            right = mid - 1  # Ищем в левой половине

    return -1  # Элемент не найден


# Пример использования
n = int(input())  # Размер массива
arr = sorted(list(map(int, input().split())))  # Отсортированный массив
target = int(input())  # Что ищем
result = binary_search(arr, target)
print(result)


#  ----------------

def binary_search_left(arr, target):
    left, right = 0, len(arr) - 1
    result = len(arr)  # Если не найдём подходящего индекса, вернём длину массива

    while left <= right:
        mid = (left + right) // 2
        if arr[mid] >= target:  # Условие: ищем первый элемент >= target
            result = mid  # Сохраняем текущий индекс как возможный результат
            right = mid - 1  # Пробуем найти ещё меньший индекс в левой половине
        else:
            left = mid + 1  # Ищем в правой половине

    return result


# Пример использования
n = int(input())  # Размер массива
arr = sorted(list(map(int, input().split())))  # Отсортированный массив
target = int(input())  # Что ищем
result = binary_search_left(arr, target)
print(result)  # Возвращает индекс первого элемента >= target

#  ---------------

# Пример: найти корень уравнения f(x) = target с точностью eps
def f(x):
    # Пример функции: x^2 - target = 0
    return x * x

def binary_search_real(target, eps=1e-6):
    left, right = 0.0, 10**9  # Задаём границы поиска
    while right - left > eps:  # Пока не достигнем нужной точности
        mid = (left + right) / 2
        if f(mid) < target:
            left = mid  # Сдвигаем левую границу
        else:
            right = mid  # Сдвигаем правую границу
    return (left + right) / 2  # Возвращаем среднее значение

# Пример использования
target = float(input())  # Например, найти корень из числа
result = binary_search_real(target)
print(f"{result:.6f}")  # Выводим с точностью 6 знаков после запятой