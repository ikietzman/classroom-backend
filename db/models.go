package db

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name          string `json:"name"`
  Username      string `json:"username";gorm:"unique;not null"`
  Email         string `json:"email"`
  Uid           string
}

type Course struct {
  gorm.Model
  Name          string
  Description   string
}

type Lesson struct {
  gorm.Model
  CourseID      uint
  Name          string
  Description   string
  Content       string
}

type UserCourse struct {
  gorm.Model
  CourseID      uint
  UserID        uint
}

type Test struct {
  gorm.Model
  CourseID      uint
  Name          string
  Description   string
}

type Question struct {
  gorm.Model
  TestID        uint
  Question      string
  Option1       string
  Option2       string
  Option3       string
  Option4       string
  Answer        string
}

type UserTest struct {
  gorm.Model
  UserID        uint
  TestID        uint
}

type UserQuestion struct {
  gorm.Model
  UserID        uint
  UserTestID    uint
  QuestionID    uint
  Answer        string
}
