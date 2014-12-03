package blog

import (
    "net/http"
    "html/template"
    "io/ioutil"
    "strings"

    "go-websiteskeleton/app/common"

    "github.com/gorilla/mux"
)

type Page struct {
    Title  string
    Active  string
    Body template.HTML
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(string("content/" + filename))
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Active: "blog_view", Body: template.HTML(string(body))}, nil
}
func GetViewPage(rw http.ResponseWriter, req *http.Request) {

    params := mux.Vars(req)
    title  := params["title"]
    p, _   := loadPage(title)

    p.Title = strings.Title(strings.Replace(title, "-", " ", -1))

    common.Templates = template.Must(template.ParseFiles("templates/blog/view.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}