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

func UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	oldname := params.ByName("name")
	r.ParseForm()
	var (
		read io.Reader = r.Body
		user model.User
	)
	// read = io.TeeReader(read, os.Stderr)
	json.NewDecoder(read).Decode(&user)

	err := updateUser(user, oldname)
	if err == nil {
		io.WriteString(w, "Done")
	} else {
		io.WriteString(w, err.Error())
	}
}

func updateUser(newUser model.User, oldname string) error {
	// var user model.User
	db := DButil.GetClient()
	res := db.Model(&model.User{}).Where("name = ?", oldname).Update(model.User{Name: newUser.Name, Password: newUser.Password})
	// user.Name = newUser.Name
	// user.Password = newUser.Password
	// res := db.Save(&user)
	return res.Error
}
