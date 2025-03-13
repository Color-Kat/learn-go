#include <iostream>
#include <vector>
#include <sstream>
#include <string>

using namespace std;

int main() {
    int n;
    cin >> n;
    cin.ignore(); // очищаем символ перевода строки после ввода n

    vector<vector<int>> numbers(n); // создаем вектор векторов для хранения всех чисел

    // Считываем n строк
    for(int i = 0; i < n; i++) {
        string line;
        getline(cin, line); // считываем всю строку

        stringstream ss(line); // создаем поток для разбора строки
        int num;

        // Извлекаем числа из строки и добавляем в вектор
        while(ss >> num) {
            numbers[i].push_back(num);
        }
    }

    // Выводим результат для проверки (можно убрать для контеста)
    for(int i = 0; i < n; i++) {
        for(int j = 0; j < numbers[i].size(); j++) {
            cout << numbers[i][j] << " ";
        }
        cout << endl;
    }

    return 0;
}