package about

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
        Active: "about",
        Title: "About",
    }

    common.Templates = template.Must(template.ParseFiles("templates/about/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}