package main

import (
    "net/http"
    "html/template"
    "fmt"

    "github.com/tst_martini/models"
    "github.com/go-martini/martini"
)

var discounts map[string]*models.Discounts

func main() {
    discounts = make(map[string]*models.Discounts,0)
    m := martini.Classic()

    staticOptions := martini.StaticOptions{Prefix:"static"}
    m.Use(martini.Static("static", staticOptions))
    m.Get("/", indexHandler)
    m.Get("/write", writeHandler)
    m.Get("/Save", saveDiscountsHandler)
    m.Get("/test", func()string{
    return "That's ok!"
    })

    m.Run()
}

func indexHandler(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
    if err != nil{
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "index", nil)
}

func writeHandler(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
    if err != nil{
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "write", nil)
}