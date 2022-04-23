package main

import (
  "net/http"
  router "github.com/iankietzman/router"
)


func main() {

  r := router.NewRouter()

	http.ListenAndServe(":3000", r)
  
}
