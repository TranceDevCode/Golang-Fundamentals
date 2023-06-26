package main

import (
	"fmt"
	"reflect"
)

// Maps, el mapa es un arreglo pero compuesto por mas de un valor, y permite almacenar en formato clave-valor, los clasicos objetos
// mas menos en Javascript es algo asi {"id": 1, "nombre": "Chile"}, se podria decir que es algo parecido
func main() {
	paises := make(map[string]int)
	paises["Argentina"] = 40
	paises["Espana"] = 46
	paises["Brasil"] = 190
	paises["Uruguay"] = 34
	paises["Chile"] = 22
	fmt.Println(paises)
	fmt.Println(reflect.TypeOf(paises))
	fmt.Println(paises["Chile"])
	fmt.Println("-----------------------Otro Mapa------------------------------------------")
	paises2 := map[int]string{
		1:  "Chile",
		2:  "Peru",
		3:  "Brasil",
		4:  "Mexico",
		5:  "Venezuela",
		6:  "Argentina",
		7:  "Espana",
		8:  "Dinamarca",
		11: "USA",
	}
	fmt.Println(paises2)
	//para acceder o mostrar Chile, accedo al id, que podria ser un id de una tabla
	fmt.Println(paises2[1])

	//Veamos si existe algun valor en el map
	/*
		Declaramos 2 variable, pais y existe
		pais, va a tomar el nombre del pais
		existe va a retornar un boolean
	*/
	fmt.Println("-----------------------IF en Map------------------------------------------")

	//De esta forma podria hacer mineria de datos dentro de un map
	pais, existe := paises2[11]
	if existe {
		fmt.Println("Si existe el pais: ", pais)
	} else {
		fmt.Println("No existe el pais: ", pais)
	}

	//Eliminar un elemento de un pais
	fmt.Println("-----------------------Eliminando un elemento en Map------------------------------------------")
	delete(paises2, 1)
	fmt.Println(paises2)

	//Recorriendo un map
	fmt.Println("-----------------------Recorriendo un map con ciclo FOR------------------------------------------")
	for id, valor := range paises2 {
		fmt.Printf("ID: %v | Nombre: %v \n", id, valor)
	}

	fmt.Println("-----------------------Mapa que usa el profesor para sus API------------------------------------------")
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Metodo POST con Maps ",
	}
	fmt.Println(respuesta)
	fmt.Println("estado=", respuesta["estado"])

}

/* //Arreglos y Slices

func main() {
	//arreglo (array)
	var paises [5]string
	paises[0] = "Chile"
	paises[1] = "Peru"
	paises[2] = "Bolivia"
	paises[3] = "Venezuela"
	paises[4] = "Espana"
	fmt.Println("-----------------------PAISES------------------------------------------")
	fmt.Println(paises)
	//asi accedemos a una posicion del arreglo
	fmt.Println(paises[2])
	//conociendo el largo de un array
	fmt.Println("El largo del arreglo es: ", len(paises))
	//consultando el tipo de dato del arreglo
	fmt.Println(reflect.TypeOf(paises))

	//Slice, un Slice es basicamente un arreglo dinamico
	fmt.Println("-----------------------SLICES------------------------------------------")
	var paises2 = make([]string, 5)
	paises2[0] = "Mexico"
	paises2[1] = "Ecuador"
	paises2[2] = "Argentina"
	paises2[3] = "Uruguay"
	paises2[4] = "Espana"
	fmt.Println(paises2)
	//asi accedemos a una posicion del arreglo
	fmt.Println(paises2[2])
	//conociendo el largo de un array
	fmt.Println("El largo del arreglo es: ", len(paises2))
	//consultando el tipo de dato del arreglo
	fmt.Println(reflect.TypeOf(paises2))

	//agregar un elemento al slice
	paises2 = append(paises2, "Noruega")
	fmt.Println("-----------------------Agregando un Elemento del slice------------------------------------------")
	fmt.Println(paises2)
	//asi accedemos a una posicion del arreglo
	fmt.Println(paises2[2])
	//conociendo el largo de un array
	fmt.Println("El largo del arreglo es: ", len(paises2))
	//consultando el tipo de dato del arreglo
	fmt.Println(reflect.TypeOf(paises2))

	//Eliminar un elemento al Slice
	paises2 = append(paises[:2], paises2[2+1:]...)
	fmt.Println("-----------------------Eliminando un Elemento del slice------------------------------------------")
	fmt.Println(paises2)
	//asi accedemos a una posicion del arreglo
	fmt.Println(paises2[1])
	//conociendo el largo de un array
	fmt.Println("El largo del arreglo es: ", len(paises2))
	//consultando el tipo de dato del arreglo
	fmt.Println(reflect.TypeOf(paises2))

} */

/*
//ciclos e iteraciones

func main() {
	i := 1
	for i < 11 {
		fmt.Println(i)
		i++
	}
	fmt.Println("-----------------------------------------------------------------")
	for i2 := 1; i2 < 11; i2++ {
		fmt.Println(i2)
	}
	fmt.Println("-----------------------------------------------------------------")
	for i3 := 0; i3 < 11; i3 += 5 {
		fmt.Println(i3)
	}
	fmt.Println("-----------------------------------------------------------------")
	for i4 := 100; i4 > 10; i4 -= 5 {
		fmt.Println(i4)
	}
	fmt.Println("-----------------------------------------------------------------")
	for i5 := 10; i5 > 0; i5-- {
		fmt.Println(i5)
	}
	fmt.Println("-----------------------------------------------------------------")
	for i6 := 1; i6 < 21; i6++ {
		if i6 == 6 {
			//break -> hasta aca va a llegar el for
			//continue -> con continue se salta el i6 ==6, no lo muestra y continua con los otros
			continue
		}
		fmt.Println(i6)
	}
	fmt.Println("-----------------------------------------------------------------")
	for i7 := 1; i7 < 11; i7++ {
		//Usando el Resto
		if i7%2 == 0 {
			fmt.Println(i7)
		}
	}
} */

/* //Condicionales If, if else, switch case, tambien denominados comparadores de

func main() {
	//Operadores de comparacion
	// x == y	Es igual x igual a y ?
	// x != y	Es x diferente de y ?
	// x < y	Es x menor que y?
	// x <= y	Es x menor o igual que y?
	// x > y	Es x mayor que y?
	// x >= y	Es x mayor o igual que y?

	edad := 11
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
	fmt.Println("-----------------------------------------------------------------")

	color := "azul"
	if color == "rojo" {
		fmt.Println("Es rojo como la sangre de los araucanos")
	} else if color == "blanco" {
		//EL else if se denomina if anidado
		fmt.Println("Es blanco como la nieve")
	} else if color == "azul" {
		fmt.Println("Es azul como el cielo")
	} else {
		fmt.Println("No tiene color patrio")
	}
	fmt.Println("-----------------------------------------------------------------")
	//Operadores logicos && (and) || (or) ! (not)
	if color == "azul" && edad == 11 {
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("color es azul y edad es 11")
	}

	//Declarar variables en una condición
	if variable := 2; variable == 1 {
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("Variable es igualitico a 2")
	}

	//Switch Case
	switch color {
	case "rojo":
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("Es rojo como la sangre de los araucanos")
		break
	case "azul":
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("Es azul como el cielo desde un switch case")
		break
	case "blanco":
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("Es blanco como la nieve")
		break
	default:
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("No tiene color patrio")
		break
	}

} */

/* // Punteros: Me permite acceder a un valor binario de el objeto, el puntero lo uso con & y basicamente me muestra el espacio o la posicion en memoria
func main() {
	var estado bool = true
	color := "rojo"
	fmt.Println(color, &color)
	fmt.Println(estado, &estado)
} */

/* //Reflect y TypeOf
func main() {
	//var string1 string = "texto"
	string1 := "algo Nandu"
	fmt.Println(reflect.TypeOf(string1))
} */

/* //Tipos de datos
func main() {
	var string1 string = "texto"
	fmt.Println(string1)
	textGrande := `lorem ipsum y blah blah eeeeh lorem ipsum y blah blah eeeeh lorem ipsum y blah blah eeeeh lorem ipsum y blah blah eeeeh lorem ipsum y blah blah eeeeh`
	fmt.Println(textGrande)
	var estado bool = true
	fmt.Println(estado)
	var flotante32 float32 = 32.33
	fmt.Println(flotante32)
	var flotante64 float64 = 64.33
	fmt.Println(flotante64)
	var entero int = 1234
	fmt.Println(entero)
	var entero8 int8 = 123 // -128 a 127
	fmt.Println(entero8)
	var entero16 int16 = 123 // -2^15 a 2^15 -1
	fmt.Println(entero16)
	var entero32 int32 = 45611 // -2^31 a 2^31 -1
	fmt.Println(entero32)
	var entero64 int64 = 78910 // -2^63 a 2^63 -1
	fmt.Println(entero64)

	var entero_unit8 uint8 = 233 // 0 a 255
	fmt.Println(entero_unit8)
	var entero_unit16 uint16 = 12451 // 0 a 2^16 -1
	fmt.Println(entero_unit16)
	var entero_unit32 uint32 = 789123235 // 0 a 2^32 -1
	fmt.Println(entero_unit32)
	var entero_unit64 uint64 = 233 // 0 a 2^64 -1
	fmt.Println(entero_unit64)
}
*/
/*

//variables y constantes

const MiConstante = "Valor de mi constante"

func main() {
	//declaracion por inferencia

	var nombre string = "Miguel"
	fmt.Println(nombre)
	//declaracion rapida o corta
	nombre2 := "Juan"
	fmt.Println(nombre2)

	fmt.Printf("El Valor de mi constante es: %s", MiConstante)

}


*/
