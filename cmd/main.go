package main

import (
	"log"
	"net/http"
)

const STATIC_DIR = "../static"

type Rules map[string] string

var rules Rules

func GetRules() Rules {
    return map[string] string{
        "google.com": "/config/config.json",
        "crumbo.biz": "/config/config.test.json",
    }
}


func handleFileRequest(w http.ResponseWriter, req *http.Request) {
    log.Printf("===Request for config file made")
    

    //Default, give em what they ask for
    configFilePath := STATIC_DIR + req.URL.Path
    log.Printf("   File requested: %v", configFilePath)

    //Actual, check no overriding rules apply, and access is permitted
    if override := rules[req.Host]; override != "" {
        configFilePath = STATIC_DIR + override
        log.Printf("   File mapped to: %v", configFilePath)
    }


    http.ServeFile(w, req, configFilePath)
}

func main() {
   
    rules = GetRules()

    log.Printf("Config server has started\n")

    http.Handle("/config/*", http.HandlerFunc(handleFileRequest))

    if err := http.ListenAndServe(":46101", nil); err != nil {
        panic("Oh no!")
    }
}
