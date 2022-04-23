package handlers

import (
  "encoding/json"
  "net/http"
  "strconv"
  "github.com/go-chi/chi/v5"
  db "github.com/iankietzman/db"
)

func PostUserQuestion(w http.ResponseWriter, r *http.Request) {
  var userQuestion db.UserQuestion
  json.NewDecoder(r.Body).Decode(&userQuestion)
  userID, _ := strconv.ParseUint(chi.URLParam(r, "user_id"), 10, 64)
  usertestID, _ := strconv.ParseUint(chi.URLParam(r, "usertest_id"), 10, 64)
  userQuestion.UserID = uint(userID)
  userQuestion.UserTestID = uint(usertestID)
  // if !ok {
  //   http.Error(w, http.StatusText(422), 422)
  //   return
  // }
  database.Create(&userQuestion)
  b, _ := json.Marshal(userQuestion)
  w.Write(b)
}

func GetUserQuestion(w http.ResponseWriter, r *http.Request) {
  var userQuestion []db.UserQuestion
  testID, _ := strconv.ParseUint(chi.URLParam(r, "usertest_id"), 10, 64)
  tID := uint(testID)
  database.Where("user_test_id = ?", tID).Find(&userQuestion)
  b, _ := json.Marshal(userQuestion)
  w.Write(b)
}
