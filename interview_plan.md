# Базовые знания ГО  
## Конкурентность  
### WaitGroup - механизм ожидания горутин  
wg.Add() добавлять перед созданием горутины  
wg.Done()  defer  
wg.Wait()  
### Data race
Data race - это когда несолько горутин обращаются к одним и тем же данным, и хотя бы одна из них записывает(изменяет) данные.
## Race candition  
Race condition — это недостаток, возникающий, когда время или порядок событий влияют на правильность программы    
Решение проблемы data race и Race candition  Mutex  
### Mutex - механизм получения эксклюзивной блокировки  
mu.Lock()  
mu.Unlock()  
### RWMutex - позволяет отдельно заблокировть чтение/запись  
Другими словами, читателям не нужно ждать друг друга. Им нужно только дождаться, пока писатели удержат замок. Можно параллельно читать, но не писать   
mu.Lock - блокировка на чтение и запись  
mu.Rlock - блокировка на чтение, могут читать несколько горутин, но нельзя писать  

### Deadlock
Deadlock возникает, когда группа goroutines ждет друг друга, и ни одна из них не может продолжить.
### Каналы   
Каналы (channels) представляют инструменты коммуникации между горутинами.  
### Небуфферизированные каналы(не указан cap при инициализации)  
Если канал пустой, то горутина-получатель блокируется, пока в канале не окажутся данные. Когда горутина-отправитель посылает данные, горутина-получатель получает эти данные и возобновляет работу.  
Горутина-отправитель может отправлять данные только в пустой канал. Горутина-отправитель блокируется до тех пор, пока данные из канала не будут получены.  
### Буферизированные каналы (указан cap при инициализации)  
Если канал пуст, то получатель ждет, пока в канале появится хотя бы один элемент.  
При отправке данных горутина-отправитель ожидает, пока в канале не освободится место для еще одного элемента и отправляет элемент, только тогда, когда в канале освобождается для него место.  
Close() закрытие канала.  
Писать в закрытый канал паника.  
Закрыть канал 2 раза паника.  
Канал возврашает два значения, само значение и флаг закрыт канал или нет.  
### Worker pool  
```
func workerPool() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}
	go func() {
		for i := 0; i < 1000; i++ {
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}
```
## ООП  
Абстракция   
Инкапсуляция про контроль доступа к свойствам объекта и их динамическая валидация/преобразования  
Наследование это возможность наследоваться одним объектам от других, "перенимая" все методы родительских объектов.   
Полиморфизм - это возможность объектов использовать один и тот же интерфейс  
## SOLID  
### S – Single Responsibility (Принцип единственной ответственности)  
Каждый класс должен отвечать только за одну операцию.  
### O — Open-Closed (Принцип открытости-закрытости)  
Классы должны  быть  открыты для расширения, но закрыты для модификации.  
### L — Liskov Substitution (Принцип подстановки Барбары Лисков)  
Если П является подтипом Т, то любые объекты типа Т, присутствующие в программе, могут заменяться объектами типа П без негативных последствий для функциональности программы.  
### I — Interface Segregation (Принцип разделения интерфейсов)  
Не следует ставить клиент в зависимость от методов, которые он не использует.  
### D — Dependency Inversion (Принцип инверсии зависимостей)  
Модули верхнего уровня не должны зависеть от модулей нижнего уровня. И те, и другие должны зависеть от абстракций. Абстракции не должны зависеть от деталей. Детали должны зависеть от абстракций.  
## Context
withCancel, cancel := contect.WithCancel(ctx)  
cancel() - вызывать на уровне создания контекста withCancel  

# Спросят про pet проект  
Пет проект musique музыкальный сервис, основанныый на микросервисной архитектуре.  https://github.com/gxrlxv/musique
# Запросы  
## SELECT FOR UPDATE

если запросить FOR UPDATE - то мы можем быть уверены, что ни одна другая транзакция не сможет обновить эту строку до конца нашей транзакции.
```
SELECT ... FROM table WHERE ... FOR UPDATE  Блокирует строки всех таблицы учствующийх в запросе  
SELECT ... FROM table WHERE ... FOR UPDATE OF list-of-tablenames Блокирует строки перечисленных таблиц  
```
## Inner join
```
SELECT o.OrderID, c.CustomerName  
FROM Orders o
INNER JOIN Customers с
ON o.CustomerID = с.CustomerID;
```
## Агрегатные функции  
## group by
```
SELECT Manufacturer, COUNT(*) AS ModelsCount  
FROM Products  
GROUP BY Manufacturer  
```
Первый столбец в выражении SELECT - Manufacturer название группы, а второй столбец - ModelsCount представляет результат функции Count, которая вычисляет количество строк в группе.  
## Having  
Для проверки условия аггрегатной функции использовать Having  
```
SELECT country, city, count(*)  
FROM customers  
GROUP BY country  
HAVING count(*) > 3  
```
##  insert update 
```
INSERT INTO the_table (id, column_1, column_2)   
VALUES (1, 'A', 'X')  
ON CONFLICT (id) DO UPDATE   
  SET column_1 = excluded.column_1,    
      column_2 = excluded.column_2;  
```
##  select delete  
```
DELETE FROM tableA WHERE entitynum IN (...your select...)  
DELETE FROM tableA WHERE (...your select...)
```
## With
```
WITH query_name1 AS (  
     SELECT ...  
     )  
   , query_name2 AS (  
     SELECT ...  
       FROM query_name1  
        ...  
     )  
SELECT ...   
FROM query_name1
```
## Union  
Объединение селектов  
```
SELECT column_name(s) FROM table1  
UNION  
SELECT column_name(s) FROM table2;  
```

# Теори бд
## Репликация   
Репликация - дублирование бд. Есть мастер в который пишут, мастер отправляет в слейв изменения. Можно синхронно и асинхронно 
## Шардирование  
Шардинг — метод разделения и хранения единого логического набора данных в виде множества баз данных.  Разбиение на бд происходит по определенным ключам   
## Транзакция  
Транзакция — это набор операций по работе с базой данных (БД), объединенных в одну атомарную пачку.

## MVCC 
Multiversion Concurrency Control, Многоверсионное управление конкурентным доступом  
Это означает, что каждый SQL-оператор видит снимок данных (версию базы данных) на определённый момент времени, вне зависимости от текущего состояния данных. Это защищает операторы от несогласованности данных, возможной, если другие конкурирующие транзакции внесут изменения в те же строки данных, и обеспечивает тем самым изоляцию транзакций для каждого сеанса баз данных. MVCC, отходя от методик блокирования, принятых в традиционных СУБД, снижает уровень конфликтов блокировок и таким образом обеспечивает более высокую производительность в многопользовательской среде.  
Основное преимущество использования модели MVCC по сравнению с блокированием заключается в том, что блокировки MVCC, полученные для чтения данных, не конфликтуют с блокировками, полученными для записи, и поэтому чтение никогда не мешает записи, а запись чтению.

## ACID  
ACID - набор требований к транзакционной системе, обеспечивающий наиболее надёжную и предсказуемую её работу — атомарность, согласованность, изоляция, устойчивость.  

Атомарность - гарантирует, что никакая транзакция не будет зафиксирована в системе частично. Будут либо выполнены все её подоперации, либо не выполнено ни одной.  

Согласованность - каждая успешная транзакция по определению фиксирует только допустимые результаты. Это больше бизнес правило, что с счетов будет снята одинаковая сумма  

Изоляция во время выполнения транзакции параллельные транзакции не должны оказывать влияния на её результат. Изолированность — требование дорогое, поэтому в реальных базах данных существуют режимы, не полностью изолирующие транзакцию

Устойчивость независимо от проблем на нижних уровнях (к примеру, обесточивание системы или сбои в оборудовании) изменения, сделанные успешно завершённой транзакцией, должны остаться сохранёнными после возвращения системы в работу


## Уровни изоляций  

Read uncommitted когда мы читаем или пишем данные, мы не блокируем другим пользователям ни чтение, ни запись этих данных. Он гарантирует, что все транзакции, которые пришли в базу данных, будут выполнены. Если два пользователя одновременно начали выполнять запросы с одними и теми же данными, то обе эти транзакции будут выполнены последовательно.  

Read committed другая транзакция никогда не видит промежуточных этапов первой транзакции. Мы гарантируем, что у нас никогда не будет ситуации, когда мы видим какие-то части данных, недописанные данные.

Repeatable read мы берем селект в нашу транзакцию, то получается как будто слепок данных. И мы в этот момент не видим изменений других пользователей, все время работаем именно с этим слепком данных. Минус в том, что мы блокируем данные и, соответственно, у нас меньше параллельных запросов, которые могут работать с данными. Могут появиться новые данные и мы их не увидим это минус.

Serializable это полная блокировка данных в таблице. Она спасает от фантомного чтения, то есть от чтения как раз тех данных, которые у нас добавились или удалились, потому что мы блокируем таблицу, не разрешаем в нее писать. 

В отличие от последовательного выполнения, сериализуемость позволяет выполнять несколько транзакций одновременно, но с одним условием: результат должен быть эквивалентен последовательному выполнению.


## Индексы  

Индексы создаются для столбцов таблиц и представлений. Индексы предоставляют путь для быстрого поиска данных на основе значений в этих столбцах. Например, если вы создадите индекс по первичному ключу, а затем будете искать строку с данными, используя значения первичного ключа, то SQL Server сначала найдет значение индекса, а затем использует индекс для быстрого нахождения всей строки с данными. Без индекса будет выполнен полный просмотр (сканирование) всех строк таблицы, что может оказать значительное влияние на производительность.  

Индекс состоит из набора страниц, узлов индекса, которые организованы в виде древовидной структуры — сбалансированного дерева. Эта структура является иерархической по своей природе и начинается с корневого узла на вершине иерархии и конечных узлов, листьев, в нижней части.  

Листья индекса могут содержать как сами данные таблицы, так и просто указатель на строки с данными в таблице, в зависимости от типа индекса: кластеризованный индекс или некластеризованный.   Кластеризованный индекс хранит реальные строки данных в листьях индекса.  
В отличие от кластеризованного индекса, листья некластеризованного индекса содержат только те столбцы (ключевые), по которым определен данный индекс, а также содержит указатель на строки с реальными данными в таблице.

Индексы в базе данных — это отдельная структура. Таблица от нее не зависима. То есть индекс вы можете в любой момент удалить и перестроить заново, и таблица от этого не пострадает. 

У нас действительно ускорятся селекты. Каждый раз, когда нам нужно пройти по какому-то значению, мы заходим в индекс, находим там ссылку на сами значения. Индексы, как правило, содержат именно ссылки на строки, а не сами строки. И для селектов это работает идеально. Но как только мы захотим задать данные таблицы, проапдейтить либо удалить данные, то все эти деревья придется перестраивать. 



# Знания по сетям и UNIX  
## Память  
https://ru.stackoverflow.com/questions/277295/Представление-кучи-и-стека#:~:text=Стек%20как%20бы%20быстрее%20потому%2C,(если%20не%20сотни)%20телодвижений%20процессора  
## Потоки  
https://habr.com/ru/post/40227/  
https://ru.wikipedia.org/wiki/Поток_выполнения  
 Поток (thread) — это, сущность операционной системы, процесс выполнения на процессоре набора инструкций, точнее говоря программного кода.  
 Процесс (process) — не что более иное, как некая абстракция, которая инкапсулирует в себе все ресурсы процесса (открытые файлы, файлы отображенные в память...) и их дескрипторы, потоки и т.д. Каждый процесс имеет как минимум один поток. Также каждый процесс имеет свое собственное виртуальное адресное пространство и контекст выполнения, а потоки одного процесса разделяют адресное пространство процесса.
# Алгоритмтическая секция  
### База:  
- Массив +  
- Множество   +
- Хэш-отображение (HashMap, хэшмэп) +  
- Связный список +  
- Стек   LIFO +
- Очередь   FIFO +
- Дерево +
- Граф
--- 
- Поиск в ширину BFS + 
  (queue.push_back(соседи) -> queue[0])
- Поиск в глубину DFS + 
  (queue.push_back(соседи) -> stack[len(stack)-1])
- Двоичный поиск +
  (дихотомия, поиск в отсортированном массиве делением пополам)
- Quickselect 
  (алгоритм нахождения k-го наименьшего элемента в неотсортированном массиве; выбираем pivot -> разделяем на элементы меньше и больше -> если k = позиции pivot возвращаем результат, если k < длины левого, то повторяем для левого, иначе для  правого)
- Алгоритм Дейкстры
  ( https://ru.wikipedia.org/wiki/Алгоритм_Дейкстры 
  Поиск кратчайшего расстояния между точками через bfs)
- Алгоритм Беллмана – Форда 
  (Допускает отрицательные веса
  https://habr.com/ru/company/otus/blog/484382/
  https://e-maxx.ru/algo/ford_bellman )
- А-звезда ( https://ru.wikipedia.org/wiki/A* )
---
- Рекурсия +
- Метод двух указателей +
- Принцип «разделяй и властвуй» +
- Метод скользящего окна +
### Некст:  
- Куча (также известная как приоритетная очередь)
- LRU-кэш (2Q)
  (хранить hashmap и список по времени запросов -> если список > N то удалить самый старый элемент)
- Двоичные деревья (AVL - баланс по высоте, красно-чёрные) +
- Двоичное дерево поиска (слева меньше, справа больше) +
- Непересекающиеся множества  
### Сортировка

- Быстрая сортировка +
  (pivot -> в конец -> меньшие влево -> same для лев и прав)
- Сортировка слиянием
  (делим пополам до единичной длины -> мержим все части в одну через два указателя)
- Топологическая сортировка
  (DFS -> обойти все ноды -> проверить маркировку ноды (временная-stop/посещенная-next) -> пометить ноду временной -> запустить обход всех смежных нод -> пометить ноду посещенной и добавить в рез)
- Сортировка подсчётом
  (записать кол-во вхождений в массив -> прибавить каждому эл-ту кол-во пред эл-та (индексы) -> обойти исходный массив, расставляя в рез по индексу и убавляя индекс на 1)
  (Применение сортировки подсчётом целесообразно лишь тогда, когда сортируемые числа имеют (или их можно отобразить в) диапазон возможных значений, который достаточно мал по сравнению с сортируемым множеством, например, миллион натуральных чисел меньших 1000)
## Динамическое программирование


https://habr.com/ru/post/191498/
https://www.geeksforgeeks.org/dynamic-programming/

Динамическое программирование пользуется следующими свойствами задачи:

- перекрывающиеся подзадачи;
- оптимальная подструктура;
- возможность запоминания решения часто встречающихся подзадач.

Динамическое программирование обычно придерживается двух подходов к решению задач:

- нисходящее динамическое программирование: задача разбивается на подзадачи меньшего размера, они решаются и затем комбинируются для решения исходной задачи. Используется запоминание для решений уже решенных подзадач.
- восходящее динамическое программирование: все подзадачи, которые впоследствии понадобятся для решения исходной задачи просчитываются заранее и затем используются для построения решения исходной задачи.

## Поиск с возвратом

backtracking — общий метод нахождения решений задачи, в которой требуется полный перебор всех возможных вариантов в некотором множестве М. Как правило, позволяет решать задачи, в которых ставятся вопросы типа: «Перечислите все возможные варианты …», «Сколько существует способов …», «Есть ли способ …», «Существует ли объект…» и т. п.

Решение задачи методом поиска с возвратом сводится к последовательному расширению частичного решения. Если на очередном шаге такое расширение провести не удается, то возвращаются к более короткому частичному решению и продолжают поиск дальше. Для ускорения метода стараются вычисления организовать таким образом, чтобы как можно раньше выявлять заведомо неподходящие варианты.

https://medium.com/nuances-of-programming/поиск-с-возвратом-в-решении-типичных-задач-на-собеседовании-fc60a32ff1a6