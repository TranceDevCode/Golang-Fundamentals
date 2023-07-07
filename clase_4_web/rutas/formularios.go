package rutas

import (
	"fmt"
	"html/template"
	"net/http"

	"clase_4_web/utilidades"
	"clase_4_web/validaciones"
)

func Formularios_get(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/formularios/formulario.html", utilidades.Fronted))
	template.Execute(response, nil)
	css_sesion, css_mensaje := utilidades.RetornarMensajeFlash(response, request)

}

func Formularios_post(response http.ResponseWriter, request *http.Request) {
	mensaje := ""
	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + " El campo nombre esta vacio."
	}
	if len(request.FormValue("correo")) == 0 {
		mensaje = mensaje + " El campo Email esta vacio."
	}
	if validaciones.Regex_correo.FindStringSubmatch(request.FormValue("correo")) == nil {
		mensaje = mensaje + " . El E-Mail ingresado no es valido "
	}
	if validaciones.ValidarPassword(request.FormValue("password")) == false {
		mensaje = mensaje + " . La contrasena debe tener al menos 1 numero, una mayuscula, y un largo de entre 6 y 20 catacteres "
	}
	//p2gHNiENUw
	if mensaje != "" {
		//fmt.Fprintln(response, mensaje)
		//return

		//Pasamos los mensajes de utilidades
		utilidades.CrearMensajeFlash(response, request, "danger", mensaje)
		//Luego hacemos redirect
		http.Redirect(response, request, "/formularios", http.StatusSeeOther)
	}
	fmt.Fprint(response, "Nombre:"+request.FormValue("nombre")+" | E-Mail: "+request.FormValue("correo")+" | Telefono: "+request.FormValue("telefono")+" | Contrasena: "+request.FormValue("password"))

}

/*
func Formularios_post(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Nombre:"+request.FormValue("nombre")+" | E-Mail: "+request.FormValue("correo")+" | Telefono: "+request.FormValue("telefono")+" | Contrasena: "+request.FormValue("password"))
} */
