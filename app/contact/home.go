package contact

import (
    "net/http"
    "html/template"
    "go-websiteskeleton/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
        Active string
    }

    p := Page{
        Active: "contact",
        Title: "Contact Us",
    }

    common.Templates = template.Must(template.ParseFiles("templates/contact/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}

func SubmitContactForm(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
        Active string
    }

    p := Page{
        Active: "contact",
        Title: "Contact Us",
    }

    msg := &common.Message{
        Email: req.FormValue("email"),
        Name: req.FormValue("name"),
        Content: req.FormValue("content"),
        Contact: req.FormValue("contactNumber"),
        Heardvia: req.FormValue("heardvia"),
        Require: req.FormValue("require"),
    }

    msg.Deliver()
    common.Templates = template.Must(template.ParseFiles("templates/contact/thanks.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}
