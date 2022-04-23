package main

import (
  "fmt"
  db "github.com/iankietzman/db"
)

func main() {
  LoadData()
}

func LoadData() {
  fmt.Println(db.DB)

  var courses = []db.Course{
    {
      Name: "Spanish Animals",
      Description: "Everybody loves los animales"},
    {
      Name: "German Numbers",
      Description: "Learn the Numbers from 0 to 10 in German"},
    {
      Name: "Portuguese Jobs",
      Description: "Here are some common job titles you might hear"},
  }

  var lessons = []db.Lesson{
    {
      CourseID: 1,
      Name: "Sección 1",
      Description: "English: the dog",
      Content: "Español: el perro"},
    {
      CourseID: 1,
      Name: "Sección 2",
      Description: "English: the cat",
      Content: "Español: el gato"},
    {
      CourseID: 1,
      Name: "Sección 3",
      Description: "English: the sloth",
      Content: "Español: el perezoso"},
    {
      CourseID: 1,
      Name: "Sección 4",
      Description: "English: the horse",
      Content: "Español: el caballo"},
    {
      CourseID: 1,
      Name: "Sección 5",
      Description: "English: the cow",
      Content: "Español: la vaca"},
    {
      CourseID: 1,
      Name: "Sección 6",
      Description: "English: the lobster",
      Content: "Español: la langosta"},
    {
      CourseID: 2,
      Name: "Kapital Zero",
      Description: "English: Zero",
      Content: "Deutsch: Zero"},
    {
      CourseID: 2,
      Name: "Kapitel Ein",
      Description: "English: One",
      Content: "Deutsch: Eins"},
    {
      CourseID: 2,
      Name: "Kapitel Zwei",
      Description: "English: Two",
      Content: "Deutsch: Zwei"},
    {
      CourseID: 2,
      Name: "Kapitel Drei",
      Description: "English: Three",
      Content: "Deutsch: Drei"},
    {
      CourseID: 2,
      Name: "Kapitel Vier",
      Description: "English: Four",
      Content: "Deutsch: Vier"},
    {
      CourseID: 2,
      Name: "Kapitel Fünf",
      Description: "English: Five",
      Content: "Deutsch: Fünf"},
    {
      CourseID: 2,
      Name: "Kapitel Sechs",
      Description: "English: Six",
      Content: "Deutsch:  Sechs"},
    {
      CourseID: 2,
      Name: "Kapitel Sieben",
      Description: "English: Seven",
      Content: "Deutsch: Sieben"},
    {
      CourseID: 2,
      Name: "Kapitel Acht",
      Description: "English: Eight",
      Content: "Deutsch:  Acht"},
    {
      CourseID: 2,
      Name: "Kapitel Neun",
      Description: "English: Nine",
      Content: "Deutsch: Neus"},
    {
      CourseID: 2,
      Name: "Kapital Zehn",
      Description: "English: Ten",
      Content: "Deutsch: Zehn"},
    {
      CourseID: 3,
      Name: "Lição Um",
      Description: "English: the firefighter",
      Content: "Português: o bombeiro / a bombeira"},
    {
      CourseID: 3,
      Name: "Lição Dois",
      Description: "English: the chef",
      Content: "Português: o cozinheiro / a cozinheira"},
    {
      CourseID: 3,
      Name: "Lição Três",
      Description: "English: the teacher",
      Content: "Português: o professor / a professora"},
    {
      CourseID: 3,
      Name: "Lição Quatro",
      Description: "English: the doctor",
      Content: "Português: a médica / o médico"},
    {
      CourseID: 3,
      Name: "Lição Cinco",
      Description: "English: the salesperson",
      Content: "Português: o vendedor / a vendedora"},
    {
      CourseID: 3,
      Name: "Lição Seis",
      Description: "English: the farmer",
      Content: "Português: a fazendeira / o fazendeiro"},
  }

  var tests = []db.Test{
    {
      CourseID: 1,
      Name: "Final Test",
      Description: "Animales españoles"},
    {
      CourseID: 2,
      Name: "Final Test",
      Description: "Deutsche Zahlen"},
    {
      CourseID: 3,
      Name: "Final Test",
      Description: "Animais Portugueses"},
  }

  // type Question struct {
  //   gorm.Model
  //   TestID        uint
  //   Question      string
  //   Option1       string
  //   Option2       string
  //   Option3       string
  //   Option4       string
  //   Answer        string
  // }

  var questions = []db.Question{
    {
      TestID: 1,
      Question: "the dog",
      Option1: "el perro",
      Option2: "el gato",
      Option3: "er perezoso",
      Option4: "el caballo",
      Answer: "el perro",},
    {
      TestID: 1,
      Question: "the cat",
      Option1: "la vaca",
      Option2: "el gato",
      Option3: "er perezoso",
      Option4: "el caballo",
      Answer: "el gato",},
    {
      TestID: 1,
      Question: "the sloth",
      Option1: "la vaca",
      Option2: "la langosta",
      Option3: "er perezoso",
      Option4: "el caballo",
      Answer: "el perezoso",},
    {
      TestID: 1,
      Question: "the horse",
      Option1: "la vaca",
      Option2: "el langosta",
      Option3: "el perezoso",
      Option4: "el caballo",
      Answer: "el caballo",},
    {
      TestID: 1,
      Question: "the cow",
      Option1: "la vaca",
      Option2: "el langosta",
      Option3: "el perezoso",
      Option4: "el perro",
      Answer: "la vaca",},
    {
      TestID: 1,
      Question: "the lobster",
      Option1: "la vaca",
      Option2: "la langosta",
      Option3: "el perezoso",
      Option4: "el perro",
      Answer: "la langosta",},
    {
      TestID: 2,
      Question: "Zero",
      Option1: "Ein",
      Option2: "Zero",
      Option3: "Zwei",
      Option4: "Vier",
      Answer: "Zero",},
    {
      TestID: 2,
      Question: "One",
      Option1: "Eins",
      Option2: "Zero",
      Option3: "Zwei",
      Option4: "Vier",
      Answer: "Eins",},
    {
      TestID: 2,
      Question: "Two",
      Option1: "Eins",
      Option2: "Zero",
      Option3: "Zwei",
      Option4: "Vier",
      Answer: "Zwei",},
    {
      TestID: 2,
      Question: "Three",
      Option1: "Drei",
      Option2: "Zwei",
      Option3: "Eins",
      Option4: "Vier",
      Answer: "Drei",},
    {
      TestID: 2,
      Question: "Four",
      Option1: "Eins",
      Option2: "Drei",
      Option3: "Zwei",
      Option4: "Vier",
      Answer: "Vier",},
    {
      TestID: 2,
      Question: "Five",
      Option1: "Drei",
      Option2: "Fünf",
      Option3: "Zwei",
      Option4: "Vier",
      Answer: "Fünf",},
    {
      TestID: 2,
      Question: "Six",
      Option1: "Drei",
      Option2: "Fünf",
      Option3: "Sechs",
      Option4: "Vier",
      Answer: "Sechs",},
    {
      TestID: 2,
      Question: "Seven",
      Option1: "Sieben",
      Option2: "Fünf",
      Option3: "Zwei",
      Option4: "Vier",
      Answer: "Sieben",},
    {
      TestID: 2,
      Question: "Eight",
      Option1: "Sieben",
      Option2: "Fünf",
      Option3: "Acht",
      Option4: "Sechs",
      Answer: "Acht",},
    {
      TestID: 2,
      Question: "Nine",
      Option1: "Sieben",
      Option2: "Fünf",
      Option3: "Neun",
      Option4: "Zehn",
      Answer: "Neun",},
    {
      TestID: 2,
      Question: "Ten",
      Option1: "Sieben",
      Option2: "Zehn",
      Option3: "Zwei",
      Option4: "Vier",
      Answer: "Zehn",},
    {
      TestID: 3,
      Question: "the firefighter",
      Option1: "o bombeiro / a bombeira",
      Option2: " professor / a professora",
      Option3: "o vendedor / a vendedora",
      Option4: "a fazendeira / o fazendeiro",
      Answer: "o bombeiro / a bombeira",},
    {
      TestID: 3,
      Question: "the chef",
      Option1: "o bombeiro / a bombeira",
      Option2: "o cozinheiro / a cozinheira",
      Option3: "o vendedor / a vendedora",
      Option4: "a fazendeira / o fazendeiro",
      Answer: "o cozinheiro / a cozinheira",},
    {
      TestID: 3,
      Question: "the teacher",
      Option1: "o bombeiro / a bombeira",
      Option2: "o professor / a professora",
      Option3: "o vendedor / a vendedora",
      Option4: "a fazendeira / o fazendeiro",
      Answer: "o professor / a professora",},
    {
      TestID: 3,
      Question: "the doctor",
      Option1: "o bombeiro / a bombeira",
      Option2: "o professor / a professora",
      Option3: "a médica / o médico",
      Option4: "a fazendeira / o fazendeiro",
      Answer: "a médica / o médico",},
    {
      TestID: 3,
      Question: "the salesperson",
      Option1: "o bombeiro / a bombeira",
      Option2: "o professor / a professora",
      Option3: "o vendedor / a vendedora",
      Option4: "a fazendeira / o fazendeiro",
      Answer: "o vendedor / a vendedora",},
    {
      TestID: 3,
      Question: "the farmer",
      Option1: "o bombeiro / a bombeira",
      Option2: "o professor / a professora",
      Option3: "o vendedor / a vendedora",
      Option4: "a fazendeira / o fazendeiro",
      Answer: "a fazendeira / o fazendeiro",},
  }

  db.DB.Create(&courses)
  db.DB.Create(&lessons)
  db.DB.Create(&tests)
  db.DB.Create(&questions)

}
