package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../dist/"))))
	http.HandleFunc("/loan/", handler)
	log.Fatal(http.ListenAndServe(":3005", nil))
}

type Loan struct {
	Ammount  float64
	Term     float64
	Interest float64
}

//loan/ammount/30000/term/30/interest/6.5
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
        fmt.Println(r.URL.Path)
        if r.URL.Path == "/index.html" {
          http.ServeFile(w, r, "index.html")
        }
	splitURL := strings.Split(r.URL.Path, "/")
        if len(splitURL) < 4 {
          return
        }
	fmt.Println(splitURL)
	ammount, err := strconv.ParseFloat(splitURL[3], 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ammount)
	term, _ := strconv.ParseFloat(splitURL[5], 64)
	interest, _ := strconv.ParseFloat(splitURL[7], 64)
	loan := Loan{ammount, term, interest}
	json, _ := json.Marshal(loan)
	w.Write(json)
}

