package user

import (
    "net/http"
    "html/template"

    "go-websiteskeleton/app/common"

    "github.com/gorilla/mux"
)

func GetViewPage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title  string
        Active  string
        UserId string
    }

    params := mux.Vars(req)
    userId := params["id"]

    p := Page{
        Active: "user_view",
        Title:  "Viewing User",
        UserId: userId,
    }

    common.Templates = template.Must(template.ParseFiles("templates/user/view.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}