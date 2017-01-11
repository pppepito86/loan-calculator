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

type DetailedResult struct {
  Month int `json:"month"`
  TotalPrincipalBefore float64 `json:"totalPrincipalBefore"`
  Interest float64 `json:"interest"`
  PrincipalPayment float64 `json:"principalPayment"`
  MonthPrincipal float64 `json:"monthPayment"`
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
        term, _ := strconv.Atoi(splitURL[5])
	interest, _ := strconv.ParseFloat(splitURL[7], 64)
        var json []byte
        if len(splitURL) == 9 {
          json = details(ammount, term, interest)
        } else {
          json = result(ammount, term, interest)
        }
	w.Write(json)
}

func result(ammount float64, term int, interest float64) []byte {
	rr := interest/1200
        A := (math.Pow(1+rr, float64(term))-1)/(rr*math.Pow(1+rr, float64(term)))
        Pv := ammount/A
	loan := Result{round(Pv), round(Pv*float64(term)), round(Pv*float64(term)-ammount)}

	json, _ := json.Marshal(loan)
        return json
}

func details(ammount float64, term int, interest float64) []byte {
  r := interest/1200
  A := (math.Pow(1+r, float64(term))-1)/(r*math.Pow(1+r, float64(term)))
  M := ammount/A
  principal := 30000.0
  drs := make([]DetailedResult, term)
  for i := 1; i <= term; i++ {
    principalPayment := principal*r
    monthPayment := 0.0
    if principal > 0 {
      monthPayment = M - principalPayment 
    }
    dr:= DetailedResult{i, principal, principalPayment, monthPayment, M}
    drs[i-1] = dr

    principal -= monthPayment
  }
  json, _ := json.Marshal(drs)
  return json
}

func round(x float64) float64 {
  return math.Floor(100 * x + 0.5) / 100
}
