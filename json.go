package main

import (
    "encoding/json"
    "net/http"
)

type Profile struct {
    Name    string   `json:"name"`
    Hobbies []string `json:"hobbies"`
}

func main() {
    http.HandleFunc("/", foo)
    http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
    profile := Profile{"SuperWang", []string{"football", "programming"}}

    js, err := json.Marshal(profile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
