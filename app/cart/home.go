package cart

import (
    "net/http"
    "html/template"

    "git-go-websiteskeleton/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
        Active string
    }

    p := Page{
        Active: "cart",
        Title: "Cart",
    }

    common.Templates = template.Must(template.ParseFiles("templates/cart/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}