package handlers

import (
  "encoding/json"
  "net/http"
  db "github.com/iankietzman/db"
)

func PostUserCourse(w http.ResponseWriter, r *http.Request) {
  var userCourse db.UserCourse
  ctx := r.Context()
  userID := ctx.Value("userID").(uint)
  courseID := ctx.Value("courseID").(uint)
  userCourse.CourseID = courseID
  userCourse.UserID = userID
  database.Create(&userCourse)
  w.Write([]byte("posted user course"))
}

func GetUserCourses(w http.ResponseWriter, r *http.Request) {
  var userCourses []db.UserCourse
  ctx := r.Context()
  userID := ctx.Value("userID").(uint)
  database.Where("user_id = ?", userID).Find(&userCourses)
  b, _ := json.Marshal(userCourses)
  w.Write(b)
}

func DeleteUserCourse(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
  courseID, ok := ctx.Value("courseID").(uint)
  if !ok {
    http.Error(w, http.StatusText(422), 422)
    return
  }
  database.Delete(&db.UserCourse{}, courseID)
  w.Write([]byte("delete user course"))
}
