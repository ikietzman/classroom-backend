package router

import (
  "time"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/cors"
  "github.com/go-chi/httprate"
  "github.com/go-chi/chi/v5/middleware"
  handlers "github.com/iankietzman/router/handlers"
)

func NewRouter() *chi.Mux {
  r := chi.NewRouter()

  // r.Use(middleware.AllowContentType("application/json"))
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(httprate.LimitByIP(100, 1*time.Minute))

  r.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

  r.Get("/courses/", handlers.GetAllCourses)
  r.Get("/courses/{course_id}", handlers.GetCourse)
  r.Get("/courses/{course_id}/lessons", handlers.GetCourseLessons)

  r.Get("/courses/tests/{course_id}", handlers.GetTests)
  r.Get("/test/{test_id}", handlers.GetTest)
  r.Get("/test/{test_id}/questions", handlers.GetTestQuestions)

  r.Get("/users/", handlers.GetAllUsers)

  r.Post("/users/", handlers.PostUser)

  r.Get("/users/id/{uid}", handlers.GetUserID)

  r.Post("/user/{user_id}/tests/{test_id}", handlers.PostUserTest)

  r.Get("/user/{user_id}/tests/{test_id}", handlers.GetUserTests)

  r.Get("/user/{user_id}/test/{usertest_id}", handlers.GetUserTest)

  r.Post("/user/{user_id}/test/{usertest_id}/question", handlers.PostUserQuestion)

  r.Get("/user/{user_id}/test/{usertest_id}/questions", handlers.GetUserQuestion)

  r.Route("/users/{id}", func(r chi.Router) {
    r.Use(UserCtx)
    r.Get("/", handlers.GetUser)
    r.Delete("/", handlers.DeleteUser)
    r.Patch("/", handlers.UpdateUser)
  })

  r.Route("/users/{id}/courses/{course_id}", func(r chi.Router) {
    r.Use(UserCourseCtx)
    r.Post("/", handlers.PostUserCourse)
    r.Get("/", handlers.GetUserCourses)
    r.Delete("/", handlers.DeleteUserCourse)

  })

  return r
}
