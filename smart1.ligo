//объявление типа numbers, который содержит пару из двух чисел (tuple)
type numbers is (int * int)

//объявление типов action, которые содержат пары чисел
type action is
| Addition of numbers
| Subtraction of numbers
| Multiplication of numbers
| Division of numbers

//объявление типа данных в хранилище смарт-контракта
type storage is int

//объявление математических функций.

//(const a : int ; const b : int) — параметры функции
//: int — тип результата функции
//is a + b — результат исполнения функции

function add (const a : int ; const b : int) : int is a + b

function subtract (const a : int ; const b : int) : int is a - b

function multiply (const a : int ; const b : int) : int is a * b

function divide (const a : int ; const b : int) : int is a / b

//объявление главной функции
//назначаем первому параметру тип action, параметру storе — тип storage
//функция возвращает данные типов list(operation) и int — пару из списка и числа
//после is идет результат исполнения функции:
//1) пустой список nil : list(operation).
//2) const result : int = — запись результата исполнения функции в константу result.
//2) case parameter of — результат исполнения объекта типа action,
//чье название совпадает с параметром входящей транзакции.
function main (const parameter : action ; const store : storage) :
  (list(operation) * int) is block {
    const result : int =
    case parameter of
  | Addition(n1, n2) -> add(n1, n2)
  | Subtraction(n1, n2) -> subtract(n1, n2)
  | Multiplication(n1, n2) -> multiply(n1, n2)
  | Division(n1, n2) -> divide(n1, n2)
  end;

  //вывод результата исполнения главной функции: пустой список операций и значение result
  } with ((nil : list(operation)), result)
