package main

//fmt es un paquete de formateo!
import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola, GOlang")

	/*
		https://go.dev/doc/tutorial/getting-started
	*/

	/*
		Variables
	*/

	var myString string = "Esto es una cadena de texto"
	fmt.Println(myString)
	myString = "Aqui he actualizado la cadena de texto"
	fmt.Println(myString)

	//myString = 6 Error
	var myString2 = "Esto es una cadena de texto del my string 2"
	fmt.Println(myString2)

	//Enteros
	var myInt int = 7
	myInt = myInt + 4
	fmt.Println(myInt)
	fmt.Println(myInt - 1)
	fmt.Println(myInt)

	fmt.Println(myString, myInt)

	//revisamos el tipo de dato de una variable
	fmt.Println(reflect.TypeOf(myInt))

	var myFloat = 6.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(float64(myInt) + myFloat)

	//true false

	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	//Variable declarada e iniciada de manera abreviada
	myString3 := "Esto es una cadena de texto string 3"
	fmt.Println(myString3)

	// Constantes

	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	//Control de flujo

	if myInt == 10 {
		fmt.Println("El Valor es 10")
	} else if myInt == 11 {
		fmt.Println("EL valor es 11")
	} else {
		fmt.Println("Es valor no es 10")
	}

	if myInt == 10 && myString == "Hola" {
		fmt.Println("El Valor es 10")
	} else if myInt == 11 || myString == "Hola" {
		fmt.Println("EL valor es 11")
	} else {
		fmt.Println("Es valor no es 10")
	}

	//Estructuras de datos

	//Array

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	//myArray[3] = 3 Error
	fmt.Println(myArray[2])
	//fmt.Println(myArray[3])

	var myArray2 [3]string
	fmt.Println(myArray2)

	//Map

	myMap := make(map[string]int)
	myMap["Miguel"] = 35
	myMap["Brais"] = 36
	myMap["Kasuma0"] = 34

	fmt.Println(myMap)
	fmt.Println(myMap["Miguel"])

	myMap2 := map[string]int{"Brais": 36, "Kasuma0": 34, "Miguel": 35}
	fmt.Println(myMap2)

	//List

	//Arego valores a una lista
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	//Bucles

	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}

	//con key trae llave valor
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	//con array trae las posiciones
	for index, value := range myArray {
		fmt.Println(index, value)
	}

	fmt.Println(myFunction())

	// Estructura en vez de clases

	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{"Miguel", 36}
	fmt.Println(myStruct)
}

//Funcion

func myFunction() string {
	fmt.Println("desde Mi Funcion")
	return "Hola desde la Function"
}
