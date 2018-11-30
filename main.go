package main

import (
    "net/http"
    "html/template"
    "fmt"

    "github.com/tst_martini/models"
    "github.com/go-martini/martini"
)

var discounts map[string]*models.Discount

func main() {
    discounts = make(map[string]*models.Discount,0)
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
    t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
    if err != nil{
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "write", nil)
}

func saveDiscountsHandler(w http.ResponseWriter, r *http.Request){
    id := r.FormValue("id")
    title := r.FormValue("title")
    description := r.FormValue("description")

    var discount *models.Discount
	if id != "" {
		discount = discounts[id]
		discount.Title = title
		discount.Description = description
	} else {
		id = GenerateId()
		discount := models.NewDiscount(id, title, description)
		discounts[discount.Id] = discount
	}

    http.Redirect(w, r, "/", 302)
}