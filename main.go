package main

import (
    "flag"
    "net/http"
    "time"

    "go-websiteskeleton/app/common"
    "go-websiteskeleton/app/home"
    "go-websiteskeleton/app/user"
    "go-websiteskeleton/app/jobs"
    "go-websiteskeleton/app/about"
    "go-websiteskeleton/app/contact"

    "github.com/golang/glog"
    "github.com/gorilla/mux"
)

func main() {
    flag.Parse()
    defer glog.Flush()

    router := mux.NewRouter()
    http.Handle("/", httpInterceptor(router))

    router.HandleFunc("/", home.GetHomePage).Methods("GET")
    router.HandleFunc("/user{_:/?}", user.GetHomePage).Methods("GET")

    router.HandleFunc("/user/view/{id:[0-9]+}", user.GetViewPage).Methods("GET")
    router.HandleFunc("/user/{id:[0-9]+}", user.GetViewPage).Methods("GET")

    router.HandleFunc("/jobs", jobs.GetHomePage).Methods("GET")
    router.HandleFunc("/about", about.GetHomePage).Methods("GET")
    router.HandleFunc("/contact", contact.GetHomePage).Methods("GET")

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    http.ListenAndServe(":8080", nil)
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
