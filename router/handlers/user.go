package handlers

import (
  "encoding/json"
  "net/http"
  "github.com/go-chi/chi/v5"
  db "github.com/iankietzman/db"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
  var users []db.User
  database.Find(&users)
  b, _ := json.Marshal(users)
  w.Write(b)
}

func GetUserID(w http.ResponseWriter, r *http.Request) {
  var user db.User
  uid := chi.URLParam(r, "uid")
  database.Where("uid = ?", uid).First(&user)
  b, _ := json.Marshal(user)
  w.Write(b)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
  user, ok := ctx.Value("user").(db.User)
  if !ok {
    http.Error(w, http.StatusText(422), 422)
    return
  }
  b, _ := json.Marshal(user)
  w.Write(b)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
  user, ok := ctx.Value("user").(db.User)
  if !ok {
    http.Error(w, http.StatusText(422), 422)
    return
  }
  database.Delete(&db.User{}, user.ID)
  w.Write([]byte("delete user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  var userData db.User
  ctx := r.Context()
  user, ok := ctx.Value("user").(db.User)
  if !ok {
    http.Error(w, http.StatusText(422), 422)
    return
  }
  json.NewDecoder(r.Body).Decode(&userData)
  database.First(&user)
  database.Model(&user).Updates(&userData)
  w.Write([]byte("updated user"))
}

func PostUser(w http.ResponseWriter, r *http.Request) {
  var user db.User
  json.NewDecoder(r.Body).Decode(&user)
  database.Create(&user)
  w.Write([]byte("posted user"))
}
