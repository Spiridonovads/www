package main

import (
  "fmt"
  "net/http"
  "html/template"
  "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func save_arcticle(w http.ResponseWriter, r *http.Request){
  log_1 := r.FormValue("log")
  pass_1 := r.FormValue("pass")

  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/vkr")
  if err != nil {
    panic(err)
  }
  defer db.Close()

  autho, err := db.Query(fmt.Sprintf("INSERT INTO `users`(`login`,`password`) VALUES('%s', '%s')", log_1, pass_1))
  if err != nil {
    panic(err)
  }
  defer autho.Close()

  http.Redirect(w, r, "/succes", http.StatusSeeOther)
}

func succes(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/succes.html")
  if err != nil {
    fmt.Fprintf(w, err.Error())
}
t.ExecuteTemplate(w, "succes", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/index.html")
  if err != nil {
    fmt.Fprintf(w, err.Error())
  }
  t.ExecuteTemplate(w, "index", nil)
}

func handleFunc() {
  //http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
  http.HandleFunc("/", index)
  http.HandleFunc("/succes", succes)
  http.HandleFunc("/save_arcticle", save_arcticle)
  http.ListenAndServe(":4000", nil)
}

func main() {
  handleFunc()
}
//handleRequest()

//db.Query(fmt.Sprintf(
//Установка данных
/*insert, err := db.Query("INSERT INTO users (login, password) VALUES (111, 25)")
if err != nil {
  panic(err)
}
defer insert.Close()

//Выборка данных
db, err := db.Query("SELECT login, password FROM users")
if err != nil {
  panic(err)
}
for res.Next() {
  var user User
  err := res.Scan(&user.Login, &user.Password)
  if err != nil {
    panic(err)
  }
  fmt.Println("jj")
} */



//func handleRequest() {
  //http.HandleFunc("/", home_page)
  //http.ListenAndServe(":4000", nil)
//}

//func home_page(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "ghbdtn")
    //tmpl, _ := template.ParseFiles("templates/home_page.html")
    //tmpl.Execute()
//}
