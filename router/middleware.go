package router

import (
  "context"
  "net/http"
  "strconv"
  "github.com/go-chi/chi/v5"
  db "github.com/iankietzman/db"
)

func CourseCtx(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var course db.Course
    courseID, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
    cID := uint(courseID)
    connection := db.Connect()
    connection.First(&course, courseID)
    ctx := context.WithValue(r.Context(), "courseID", cID)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

func UserCourseCtx(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userID, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
    courseID, _ := strconv.ParseUint(chi.URLParam(r, "course_id"), 10, 64)
    uID := uint(userID)
    cID := uint(courseID)
    ctx := context.WithValue(r.Context(), "userID", uID)
    ctx = context.WithValue(ctx, "courseID", cID)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

func UserCtx(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user db.User
    id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
    connection := db.Connect()
    connection.First(&user, id)
    ctx := context.WithValue(r.Context(), "user", user)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}
