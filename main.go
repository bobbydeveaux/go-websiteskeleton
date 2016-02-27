package main

import (
    "flag"
    "net/http"
    "time"

    "go-websiteskeleton/app/common"
    "go-websiteskeleton/app/home"
    "go-websiteskeleton/app/user"
    "go-websiteskeleton/app/contact"
    "go-websiteskeleton/app/blog"

    "github.com/golang/glog"
    "github.com/gorilla/mux"
)

func main() {
    flag.Parse()
    defer glog.Flush()

    router := mux.NewRouter()
    http.Handle("/", httpInterceptor(router))

    router.HandleFunc("/", home.GetHomePage).Methods("GET")
    router.HandleFunc("/users{_:/?}", user.GetHomePage).Methods("GET")

    router.HandleFunc("/user/view/{id:[0-9]+}", user.GetViewPage).Methods("GET")
    router.HandleFunc("/user/{id:[0-9]+}", user.GetViewPage).Methods("GET")

    router.HandleFunc("/contact", contact.GetHomePage).Methods("GET")
    router.HandleFunc("/contact", contact.SubmitContactForm).Methods("POST")

    router.HandleFunc("/blog/{title:[a-z0-9-]+}", blog.GetViewPage).Methods("GET")

    fileServer := http.StripPrefix("/dist/", http.FileServer(http.Dir("web/dist")))
    http.Handle("/dist/", fileServer)

    http.ListenAndServe(":8181", nil)
}

func httpInterceptor(router http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        startTime := time.Now()

        router.ServeHTTP(w, req)

        finishTime := time.Now()
        elapsedTime := finishTime.Sub(startTime)

        switch req.Method {
        case "GET":
            // We may not always want to StatusOK, but for the sake of
            // this example we will
            common.LogAccess(w, req, elapsedTime)
        case "POST":
            // here we might use http.StatusCreated
        }

    })
}
