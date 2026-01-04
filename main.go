package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"adventofcode/solutions"
	"adventofcode/solutions/shared"
)

var (
	yearFlag = flag.Int("year", 0, "Year of the Advent of Code puzzle (e.g., 2023)")
	dayFlag  = flag.Int("day", 0, "Day of the Advent of Code puzzle (1-25)")
	webFlag  = flag.Bool("web", false, "Start the web interface")
)

func main() {
	flag.Parse()

	if *webFlag {
		startWebServer()
		return
	}

	if *yearFlag == 0 || *dayFlag == 0 {
		fmt.Println("Usage: go run main.go -year <year> -day <day>")
		fmt.Println("       go run main.go -web")
		flag.PrintDefaults()
		os.Exit(1)
	}

	yearStr := fmt.Sprintf("%d", *yearFlag)
	dayStr := fmt.Sprintf("%02d", *dayFlag)
	solver, ok := solutions.GetSolver(yearStr, dayStr)
	if !ok {
		log.Fatalf("Solution for day %s in year %s not found.", dayStr, yearStr)
	}
	if err := shared.Run(solver, os.Stdin, os.Stdout); err != nil {
		log.Fatalf("Error running solution: %v", err)
	}
}

func startWebServer() {
	fmt.Println("Starting web server on :8080...")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, filepath.Join("web", "index.html"))
			return
		}
		if strings.HasPrefix(r.URL.Path, "/static/") {
			http.StripPrefix("/static/", http.FileServer(http.Dir("web"))).ServeHTTP(w, r)
			return
		}
		http.NotFound(w, r)
	})

	mux.HandleFunc("GET /api/years", getYearsHandler)
	mux.HandleFunc("GET /api/years/{year}/days", getDaysHandler)
	mux.HandleFunc("POST /api/years/{year}/days/{day}", solveProblemHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getYearsHandler(w http.ResponseWriter, r *http.Request) {
	years := solutions.AvailableYears()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Data any `json:"data"`
	}{
		Data: years,
	})
}

func getDaysHandler(w http.ResponseWriter, r *http.Request) {
	days := solutions.AvailableDays(r.PathValue("year"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Data any `json:"data"`
	}{
		Data: days,
	})
}

func solveProblemHandler(w http.ResponseWriter, r *http.Request) {
	year := r.PathValue("year")
	day := r.PathValue("day")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	bodyReader := strings.NewReader(string(body))

	solver, ok := solutions.GetSolver(year, day)
	if !ok {
		http.Error(w, fmt.Sprintf("Solution for day %s in year %s not found.", day, year), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	shared.Run(solver, bodyReader, w)
}
