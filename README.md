## arabic_roman_calculator
#Описание задачи
Создай консольное приложение «Калькулятор». Приложение должно читать из консоли введенные пользователем строки, числа, арифметические операции, проводимые между ними, и выводить в консоль результат их выполнения.
Калькулятор можно реализовать обычными функциями либо использовать структуру с методами, здесь это не принципиально.
##Требования:
  Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.
  Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.
  Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.
  Калькулятор умеет работать только с целыми числами.
  Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен выдать панику и прекратить работу.
  При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.
  При вводе пользователем не подходящих чисел приложение выдаёт панику и завершает работу.
  При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выдаёт панику и завершает работу.
  Результатом операции деления является целое число, остаток отбрасывается.
  Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль.     Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна выдать панику.
##Пример работы программы:
Input:
1 + 2
Output:
3

Input:
VI / III
Output:
II

Input:
I - II
Output:
Выдача паники, так как в римской системе нет отрицательных чисел.

Input:
I + 1
Output:
Выдача паники, так как используются одновременно разные системы счисления.

Input:
1
Output:
Выдача паники, так как строка не является математической операцией.

Input:
1 + 2 + 3
Output:
Выдача паники, так как формат математической операции не удовлетворяет зад
