package user

import (
	"net/http"
	"server_mod/handlers"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/user"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.Getlist)
	router.GET(userURL, h.GetUserId)
	router.PUT(userURL, h.PutUser)
	router.POST(userURL, h.PostUser)
	router.DELETE(userURL, h.DeleteUserId)
}

func (h *handler) Getlist(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("users"))
}

func (h *handler) GetUserId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("users"))
}
func (h *handler) PutUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("users"))
}
func (h *handler) PostUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("users"))
}

func (h *handler) DeleteUserId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("users"))
}
