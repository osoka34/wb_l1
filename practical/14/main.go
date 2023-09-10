package main

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func main() {

	val := make(chan int)
	//val := 1
	//val := "string"
	//val := true

	valInterface := interface{}(val)

	switch valInterface.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case bool:
		println("bool")
	case chan int:
		println("chan int")
	}
}
