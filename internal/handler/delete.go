package handler

import (
	"io"
	"net/http"

	"CRUD/model"
	"CRUD/pkg/DButil"

	"github.com/julienschmidt/httprouter"
)

func DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")

	err := deleteUser(name)
	if err == nil {
		io.WriteString(w, "Done")
	} else {
		io.WriteString(w, err.Error())
	}
}
func deleteUser(name string) error { //nameはunique (cf. model)
	db := DButil.GetClient()
	res := db.Where("name = ?", name).Delete(&model.User{})
	return res.Error
}
