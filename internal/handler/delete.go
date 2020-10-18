package handler

import (
	"io"
	"net/http"

	"CRUD/model"
	"CRUD/pkg/DButil"

	"github.com/julienschmidt/httprouter"
)

func DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// id := r.Form["id"]
	// password := r.Form["pass"]
	name := "hoge"
	password := "fuga"
	user := model.User{Name: name, Password: password}

	err := deleteUser(user)
	if err == nil {
		io.WriteString(w, "Done")
	} else {
		io.WriteString(w, "Error :"+err.Error())
	}
}
func deleteUser(user model.User) error {
	db := DButil.GetClient()
	res := db.Delete(&user)
	return res.Error
}
