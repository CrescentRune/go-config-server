package main

import (
    "net/http"
    "log"
)

const STATIC_DIR = "../static"


func handleFileRequest(w http.ResponseWriter, req *http.Request) {
    log.Printf("===Request for config file made")
    

    //Default, give em what they ask for
    configFilePath := STATIC_DIR + req.URL.Path
    log.Printf("   File requested: %v", configFilePath)

    //Actual, check no overriding rules apply, and access is permitted

    http.ServeFile(w, req, configFilePath)
}

func main() {
    log.Printf("Config server has started\n")

    http.Handle("/config/*", http.HandlerFunc(handleFileRequest))

    if err := http.ListenAndServe(":46101", nil); err != nil {
        panic("Oh no!")
    }
}
