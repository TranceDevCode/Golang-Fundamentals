package utilidades

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Fronted string = "templates/layout/frontend.html"

var Store = sessions.NewCookieStore([]byte("session-name"))

func RetornarMensajeFlash(response http.ResponseWriter, request *http.Request) {

}

func CrearMensajeFash(response http.ResponseWriter, request *http.Request) {
	session, err := Store.Get(request, "flash-session")
}
