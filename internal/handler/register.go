package handler

import (
	"io"
	"net/http"

	"CRUD/model"
	"CRUD/pkg/DButil"

	"github.com/julienschmidt/httprouter"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	name := r.Form.Get("name")
	password := r.Form.Get("password")
	user := model.User{Name: name, Password: password}

	err := createNewUser(user)
	if err == nil {
		io.WriteString(w, "Done")
	} else {
		io.WriteString(w, err.Error())
	}
}

func createNewUser(user model.User) error {
	db := DButil.GetClient()
	res := db.Create(&user)
	return res.Error
}
