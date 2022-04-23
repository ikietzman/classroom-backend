module main

go 1.17

require github.com/iankietzman/router v0.0.0-00010101000000-000000000000

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/go-chi/cors v1.2.0 // indirect
	github.com/go-chi/httplog v0.2.4 // indirect
	github.com/go-chi/httprate v0.5.3 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/iankietzman/db v0.0.0-00010101000000-000000000000 // indirect
	github.com/iankietzman/router/handlers v0.0.0-00010101000000-000000000000 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/rs/zerolog v1.18.1-0.20200514152719-663cbb4c8469 // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
	gorm.io/gorm v1.23.4 // indirect
)

replace github.com/iankietzman/db => ./db

replace github.com/iankietzman/router => ./router

replace github.com/iankietzman/router/handlers => ./router/handlers
