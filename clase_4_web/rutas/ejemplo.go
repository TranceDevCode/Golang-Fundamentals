package rutas

import (
	"html/template"
	"net/http"

	"clase_4_web/utilidades"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/home.html", utilidades.Fronted))
	template.Execute(response, nil)

	/*
		template, err := template.ParseFiles("templates/ejemplo/home.html", "templates/layout/frontend.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}
	*/
}

func Pagina404(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/404.html", utilidades.Fronted))
	template.Execute(response, nil)
}

/* func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hola Mundo desde Golang con Mux")
} */

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/nosotros.html", utilidades.Fronted))
	template.Execute(response, nil)
	/* template, err := template.ParseFiles("templates/ejemplo/nosotros.html", "templates/layout/frontend.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	} */
}

/* func Nosotros(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Algo desde la terminal")
	fmt.Fprintln(response, "Pagina sobre nosotros con Mux en Golang")
} */

func Parametros(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/parametros.html", utilidades.Fronted))
	vars := mux.Vars(request)
	//En el caso de usar gorilla-mux se recomienda para estos casos crear un mapa
	texto := "The Midnight retrowave"
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
		"texto": texto,
	}
	template.Execute(response, data)

	/* template, err := template.ParseFiles("templates/ejemplo/parametros.html", "templates/layout/frontend.html")
	vars := mux.Vars(request)
	//En el caso de usar gorilla-mux se recomienda para estos casos crear un mapa
	texto := "The Midnight retrowave"
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
		"texto": texto,
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	} */
}

/* func Parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ID = "+vars["id"]+" | SLUG="+vars["slug"])
} */

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/parametros_querystring.html", utilidades.Fronted))
	data := map[string]string{
		"id":   request.URL.Query().Get("id"),
		"slug": request.URL.Query().Get("slug"),
	}
	template.Execute(response, data)

	/*
		//La ruta la pasamos asi: http://localhost:8080/parametros-querystring?id=111&slug=mi-slug
		template, err := template.ParseFiles("templates/ejemplo/parametros_querystring.html", "templates/layout/frontend.html")
		//En el caso de usar gorilla-mux se recomienda para estos casos crear un mapa
		data := map[string]string{
			"id":   request.URL.Query().Get("id"),
			"slug": request.URL.Query().Get("slug"),
		}
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, data)
		} */
}

/* func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, request.URL)
	fmt.Fprintln(response, request.URL.RawQuery)
	fmt.Fprintln(response, request.URL.Query())
	fmt.Fprintln(response, request.URL.Query().Get("id"))
	fmt.Fprintln(response, request.URL.Query().Get("slug"))
	id := request.URL.Query().Get("id")
	slug := request.URL.Query().Get("slug")
	fmt.Println("--------------------------------------------------------")
	fmt.Fprintln(response, "id=%s | slug =%s", id, slug)
} */

type Habilidad struct {
	Nombre string
}
type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

func Estructuras(response http.ResponseWriter, request *http.Request) {

	template := template.Must(template.ParseFiles("templates/ejemplo/estructuras.html", utilidades.Fronted))
	habilidad1 := Habilidad{"Inteligencia"}
	habilidad2 := Habilidad{"VideoJuegos"}
	habilidad3 := Habilidad{"Programacion"}
	habilidad4 := Habilidad{"Deportes"}
	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3, habilidad4}
	template.Execute(response, Datos{"Miguel", 35, 1, habilidades})
	/*
		template, _ := template.ParseFiles("templates/ejemplo/estructuras.html", "templates/layout/frontend.html")
		habilidad1 := Habilidad{"Inteligencia"}
		habilidad2 := Habilidad{"VideoJuegos"}
		habilidad3 := Habilidad{"Programacion"}
		habilidad4 := Habilidad{"Deportes"}
		habilidades := []Habilidad{habilidad1, habilidad2, habilidad3, habilidad4}
		template.Execute(response, Datos{"Miguel", 35, 1, habilidades})
	*/

	/*
		template, err := template.ParseFiles("templates/ejemplo/estructuras.html")
		if err != nil {
			panic(err)
		} else {

			template.Execute(response, Datos{"Miguel", 35, 1})
		}
	*/
}
