package main

//Инкапсуляция - сокрытие данынх от пользователя, прячем детали реализации таким образом

type SomeInterface interface {
	SomeMethod(in string) (out string)
}

type someStruct struct {
	prefix string
}

func (s someStruct) SomeMethod(in string) (out string) {
	return s.prefix + in
}

func NewSomeInterface(prefix string) SomeInterface {
	return someStruct{prefix: prefix}
}
