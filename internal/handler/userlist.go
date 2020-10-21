package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"CRUD/model"
	"CRUD/pkg/DButil"

	"github.com/julienschmidt/httprouter"
	// "os"
)

func GetUserlist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allUsers, err := getUserlist()
	if err == nil {
		allUsersJSON, err2 := json.Marshal(allUsers) //err2マジ？
		if err2 != nil {
			io.WriteString(w, err2.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(allUsersJSON))
	} else {
		io.WriteString(w, err.Error())
	}
}
func getUserlist() ([]model.User, error) {
	db := DButil.GetClient()
	users := []model.User{}
	res := db.Find(&users)
	return users, res.Error
}
