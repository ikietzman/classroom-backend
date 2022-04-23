package db

import (
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

var DB = Connect()

func Connect() *gorm.DB {
  dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?charset=utf8mb4&parseTime=True&loc=Local", Credentials.Username, Credentials.Password, Credentials.Database)
  connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }
  fmt.Print("connected")

  connection.AutoMigrate(
    &User{},
    &Course{},
    &Lesson{},
    &UserCourse{},
    &Test{},
    &Question{},
    &UserTest{},
    &UserQuestion{},
  )

  c, err := connection.DB()

  // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
  c.SetMaxIdleConns(10)

  // SetMaxOpenConns sets the maximum number of open connections to the database.
  c.SetMaxOpenConns(100)

  return connection
}
