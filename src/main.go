package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
        "math"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../dist/"))))
	http.HandleFunc("/loan/", handler)
	log.Fatal(http.ListenAndServe(":3005", nil))
}

type Loan struct {
	Ammount  float64 `json:"ammount"`
	Term     float64
	Interest float64
}

type Result struct {
  MonthPayment float64 `json:"monthPayment"`
  TotalPayment float64 `json:"totalPayment"`
  InterestPayment float64 `json:"interestPayment"`
}

//loan/ammount/30000/term/30/interest/6.5
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/index.html" {
          http.ServeFile(w, r, "index.html")
        }
	splitURL := strings.Split(r.URL.Path, "/")
        if len(splitURL) < 4 {
          return
        }
	ammount, _ := strconv.ParseFloat(splitURL[3], 64)
	term, _ := strconv.ParseFloat(splitURL[5], 64)
	interest, _ := strconv.ParseFloat(splitURL[7], 64)

	rr := interest/1200
        A := (math.Pow(1+rr, term)-1)/(rr*math.Pow(1+rr, term))
        Pv := ammount/A
	loan := Result{round(Pv), round(Pv*term), round(Pv*term-ammount)}

	json, _ := json.Marshal(loan)
	w.Write(json)
}

func round(x float64) float64 {
  return math.Floor(100 * x + 0.5) / 100
}
