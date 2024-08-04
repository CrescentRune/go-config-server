package main

import (
	"log"
	"net/http"
	"regexp"
)

const STATIC_DIR = "../static"

type Rules map[string] string

var rules Rules

func GetRules() Rules {
    return map[string] string{
        "meep[ea]meep": "/config/config.json",
        "crumbo.biz": "/config/config.test.json",
    }
}


func handleFileRequest(w http.ResponseWriter, req *http.Request) {
    log.Printf("===Request for config file made")
    

    //Default, give em what they ask for
    configFilePath := STATIC_DIR + req.URL.Path
    log.Printf("   File requested: %v", configFilePath)

    //Actual, check no overriding rules apply, and access is permitted
    for key, val := range rules {
        if res, err := regexp.MatchString(key, req.Host); err != nil {
            //Unreachable due to checking when loading rules
            panic("Should not occur")
        } else if res {
            configFilePath = STATIC_DIR + val
            log.Printf("   File mapped to: %v", configFilePath)
        }
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
