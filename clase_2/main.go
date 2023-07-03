package main

// modulo personalizado
import (
	ExampleFile "clase_2/modulo_ejemplo"
	"fmt"
)

func main() {
	fmt.Println(ExampleFile.Ejemplo1())
	fmt.Println(ExampleFile.Ejemplo2("Miguel"))
}

/*
//Logs en Golang

import (
	"log"
	"os"
)

func main() {
	//err := errors.New("Este es un error de prueba")
	//log.Fatal(err)
	//fmt.Println("continuo despues del error")
	//log.Println("Starting the application...")
	//log.Fatalln("Fatal: Fatal error signal...") //este Fatal si te detiene la ejecucion de la aplicacion o script
	//log.Fatalf("%s Fatal: Fatal error signal", "Error Type")
	//log.Panic("Error: Panic Error Message")
	//log.Panicln("Error: Panic Error Message")
	//error := "todo esta mal"
	//log.Panicf("error %s", error)

	//Guardar los errores en un archivo con los permisos 0666
	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		//nil es el equivalente a null en golang
		log.Fatal(err)
	}
	//y cerrar cuando termine la ejecion
	defer f.Close()
	log.SetOutput(f)
	log.Printf("Error linea %v", 1)

} */

//Modulo OS | Argumentos

/* import (
	"flag"
	"fmt"
)


func main() {
	//go run main.go -nombre "Miguel" -edad 22
	nombre := flag.String("nombre", "", "El nombre de la persona")
	edad := flag.Int("edad", 18, "La edad de la persona")

	flag.Parse()
	fmt.Println("Tu nombre es: ", *nombre)
	fmt.Println("Tu edad es: ", *edad)

} */

/*
//Modulo Math/Rand

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	aleatorio := rand.Intn(101)
	//Mostramos numeros random desde 0 hasta 101
	fmt.Println(aleatorio)
	min := 1000
	max := 10000
	rand.Seed(time.Now().UnixNano()) //obtener la fecha en timestamp
	aleatorio2 := rand.Intn(max-min) + min
	fmt.Println(aleatorio2)

	fmt.Println("---------Prueba de generacion de password--------------")
	password := generatePassword(20, 1, 1, 1)
	fmt.Println(password)
}

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
} */

//Modulo Strings
/*
import (
	"fmt"
	"strings"
)

func main() {
	cadena := "Mi muneca me hablo"
	fmt.Println(cadena)
	fmt.Println(strings.ToUpper(cadena))
	fmt.Println(strings.ToLower(cadena))
	//
	letras := strings.Split(cadena, "")
	fmt.Println(letras)
	fmt.Println("---------Buscando la posicion en una cadena de texto--------------")
	pos := strings.Index(cadena, "muneca")
	if pos == -1 {
		fmt.Println("La palabra no esta en la cadena ", cadena)
	} else {
		fmt.Println("La palabra Mi esta contenida en", cadena, "a partir de la posicion", pos)
	}
	fmt.Println("---------Repetir una cadena--------------")
	repedita := strings.Repeat(cadena, 2)
	fmt.Println(repedita)

	fmt.Println("---------Reemplazar una palabra en una cadena--------------")
	cadenanueva := strings.Replace(cadena, "Mi", "Mi querida", -1)
	fmt.Println(cadenanueva)

	fmt.Println("---------Buscar una cadena de texto en cadenanueva e imprimirla--------------")
	//imprimimos desde la posicion 0 a la 5 dentro de una cadena
	fmt.Println(string(cadenanueva[5:9]))
} */

//Modulo Time
/* import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("la hora actual es: ", time.Now())
	fecha := time.Now()
	fmt.Println("El anio es: ", fecha.Year())
	fmt.Println("El mes es: ", fecha.Month())
	fmt.Println("El numero del mes es: ", int(fecha.Month()))
	fmt.Println("El dia es: ", fecha.Day())
	fmt.Println("El hora es: ", fecha.Hour())
	fmt.Println("El minuto es: ", fecha.Minute())
	fmt.Println("El segundo es: ", fecha.Second())
	fmt.Printf("El tipo es: %T \n", fecha)
	fmt.Printf("%v/%v/%v \n", fecha.Day(), int(fecha.Month()), fecha.Year())
	fmt.Println("---------------------------------Formato completo de fecha-----------------------------")
	fmt.Printf("%v %v de %v de %v a las %v:%v:%v \n", fecha.Weekday(), fecha.Day(),
		fecha.Month(), fecha.Year(), fecha.Hour(), fecha.Minute(), fecha.Second())

	fmt.Println("---------------------------------Operaciones con fecha-----------------------------")
	ahora := time.Now()
	fmt.Println("Fecha en este momento: ")
	fmt.Printf("%v/%v/%v \n", ahora.Day(), int(ahora.Month()), ahora.Year())
	fmt.Println("Fecha mas 20 dias: ")
	fecha1 := ahora.Add(time.Hour * 24 * 20)
	fmt.Printf("%v/%v/%v \n", fecha1.Day(), int(fecha1.Month()), fecha1.Year())
	fmt.Println("Fecha menos 20 dias: ")
	fecha2 := ahora.Add((time.Hour * 24) * -20)
	fmt.Printf("%v/%v/%v \n", fecha2.Day(), int(fecha2.Month()), fecha2.Year())
	fmt.Println("Fecha dentro de 1 anio: ")
	fecha3 := ahora.Add(365 * 24 * time.Hour)
	fmt.Printf("%v/%v/%v \n", fecha3.Day(), int(fecha3.Month()), fecha3.Year())
	fmt.Println("---------------------------------Fecha retornada y formateada desde una funcion-----------------------------")
	fmt.Println(FormatearFecha(ahora))
}

func FormatearFecha(fecha time.Time) string {
	v := fmt.Sprintf("%v/%v/%v", fecha.Day(), int(fecha.Month()), fecha.Year())
	return v
}
*/
