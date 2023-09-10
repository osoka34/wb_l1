package main

type SomeInterface interface {
	SomeMethod(in string) (out string)
}

type SomeOtherInterface interface {
	SomeInterface
	SomeOtherMethod(in string) (out string)
}

//Композиция - объединение нескольких интерфейсов в один
//По сути это псевдонаследование, но другого нет

//Нельзя одинаковые методы у интерфейсов, даже при разных параметрах
