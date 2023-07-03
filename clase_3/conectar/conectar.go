package conectar

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

//Funcion para conectarnos a la DB

var Db *sql.DB

func Conectar() {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	conection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	Db = conection
}

func CerrarConexion() {
	Db.Close()
}

//github.com/go-sql-driver/mysql
//github.com/joho/godotenv

/*
CREATE TABLE `clientes` (
`id` int NOT NULL AUTO_INCREMENT,
`nombre` varchar(100) NOT NULL,
`correo` varchar(100) NOT NULL,
`telefono` varchar(20) NOT NULL,
`fecha` datetime NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
*/
