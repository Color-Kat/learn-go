def trap(height):
    if not height:
        return 0

    # Инициализируем два указателя
    left = 0
    right = len(height) - 1
    # Переменные для хранения максимальных высот слева и справа
    left_max = right_max = 0
    # Переменная для хранения общего объёма воды
    water = 0

    while left < right:
        # Обновляем максимальные высоты
        left_max = max(left_max, height[left])
        right_max = max(right_max, height[right])

        # Если максимальная высота слева меньше, чем справа
        if left_max < right_max:
            # Добавляем воду на текущей позиции и двигаем левый указатель
            water += left_max - height[left]
            left += 1
        else:
            # Добавляем воду на текущей позиции и двигаем правый указатель
            water += right_max - height[right]
            right -= 1

    return water


# Пример использования
height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
result = trap(height)
print(f"Общий объём воды: {result}")

# Вывод:
# Общий объём воды: 6