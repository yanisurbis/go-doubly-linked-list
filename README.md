# Doubly linked list​

- run tests with `go test -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html`

```
List // тип контейнер
Len() // длинна списка
First() // первый Item 
Last() // последний Item 
PushFront(v interface{}) // добавить значение в начало 
PushBack(v interface{}) // добавить значение в конец 
Remove(i Item) // удалить элемент
```

```
​Item // элемент списка 
Value() interface{} // возвращает значение 
Next() *Item // следующий Item 
Prev() *Item // предыдущий
```
