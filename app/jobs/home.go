package jobs

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
        Active: "jobs",
        Title: "JObs",
    }

    common.Templates = template.Must(template.ParseFiles("templates/jobs/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}