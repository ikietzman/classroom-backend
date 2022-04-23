package handlers

import (
  "encoding/json"
  "net/http"
  "strconv"
  "github.com/go-chi/chi/v5"
  db "github.com/iankietzman/db"
)

func GetTests(w http.ResponseWriter, r *http.Request) {
  var test []db.Test
  courseID, _ := strconv.ParseUint(chi.URLParam(r, "course_id"), 10, 64)
  cID := uint(courseID)
  database.Where("course_id = ?", cID).First(&test)
  b, _ := json.Marshal(test)
  w.Write(b)
}

func GetTest(w http.ResponseWriter, r *http.Request) {
  var test db.Test
  testID, _ := strconv.ParseUint(chi.URLParam(r, "test_id"), 10, 64)
  tID := uint(testID)
  database.First(&test, tID)
  b, _ := json.Marshal(test)
  w.Write(b)
}

func GetTestQuestions(w http.ResponseWriter, r *http.Request) {
  var question []db.Question
  testID, _ := strconv.ParseUint(chi.URLParam(r, "test_id"), 10, 64)
  tID := uint(testID)
  database.Where("test_id = ?", tID).Find(&question)
  b, _ := json.Marshal(question)
  w.Write(b)
}
