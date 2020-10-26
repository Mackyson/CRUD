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

func CreateNewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	var (
		read io.Reader = r.Body
		user model.User
	)
	// read = io.TeeReader(read, os.Stderr)
	json.NewDecoder(read).Decode(&user)

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
