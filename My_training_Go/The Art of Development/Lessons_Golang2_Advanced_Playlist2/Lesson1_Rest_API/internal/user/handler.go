package user

import (
	"My_training_Go/My_training_Go/The Art of Development/Lessons_Golang2_Advanced_Playlist2/Lesson1_Rest_API/internal/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


var _ handlers.Handler =&handler{} // подсказка

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}

}
func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(usersURL, h.UpdateUser)
	router.PATCH(usersURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}
func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte, ("this is list of users"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte, ("this is create user"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte, ("this is user by uuid"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte, ("this is update users"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte, ("this is partially update users"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte, ("this is delete users"))
}
