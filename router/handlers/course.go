package handlers

import (
  "encoding/json"
  "net/http"
  "strconv"
  "github.com/go-chi/chi/v5"
  db "github.com/iankietzman/db"
)

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
  var courses []db.Course
  database.Find(&courses)
  b, _ := json.Marshal(courses)
  w.Write(b)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
  var course db.Course
  courseID, _ := strconv.ParseUint(chi.URLParam(r, "course_id"), 10, 64)
  cID := uint(courseID)
  database.First(&course, cID)
  b, _ := json.Marshal(course)
  w.Write(b)
}

func GetCourseLessons(w http.ResponseWriter, r *http.Request) {
  var lessons []db.Lesson
  courseID, _ := strconv.ParseUint(chi.URLParam(r, "course_id"), 10, 64)
  cID := uint(courseID)
  database.Where("course_id = ?", cID).Find(&lessons)
  b, _ := json.Marshal(lessons)
  w.Write(b)
}
