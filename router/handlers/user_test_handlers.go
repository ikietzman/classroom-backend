package handlers

import (
  "encoding/json"
  "net/http"
  "strconv"
  "github.com/go-chi/chi/v5"
  db "github.com/iankietzman/db"
)

func PostUserTest(w http.ResponseWriter, r *http.Request) {
  var userTest db.UserTest
  userID, _ := strconv.ParseUint(chi.URLParam(r, "user_id"), 10, 64)
  testID, _ := strconv.ParseUint(chi.URLParam(r, "test_id"), 10, 64)
  userTest.UserID = uint(userID)
  userTest.TestID = uint(testID)
  // if !ok {
  //   http.Error(w, http.StatusText(422), 422)
  //   return
  // }
  database.Create(&userTest)
  b, _ := json.Marshal(userTest)
  w.Write(b)
}

func GetUserTests(w http.ResponseWriter, r *http.Request) {
  var test []db.UserTest
  userID, _ := strconv.ParseUint(chi.URLParam(r, "user_id"), 10, 64)
  testID, _ := strconv.ParseUint(chi.URLParam(r, "test_id"), 10, 64)
  uID := uint(userID)
  tID := uint(testID)
  database.Where("user_id = ? AND test_id = ?", uID, tID).Find(&test)
  b, _ := json.Marshal(test)
  w.Write(b)
}

func GetUserTest(w http.ResponseWriter, r *http.Request) {
  var test db.UserTest
  testID, _ := strconv.ParseUint(chi.URLParam(r, "usertest_id"), 10, 64)
  tID := uint(testID)
  database.Find(&test, tID)
  b, _ := json.Marshal(test)
  w.Write(b)
}
