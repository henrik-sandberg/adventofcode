package main

import (
	"adventofcode/shared"
	"adventofcode/solutions/y2022"
	"adventofcode/solutions/y2023"
	"adventofcode/solutions/y2024"
	"adventofcode/solutions/y2025"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	yearFlag = flag.Int("year", 0, "Year of the Advent of Code puzzle (e.g., 2023)")
	dayFlag  = flag.Int("day", 0, "Day of the Advent of Code puzzle (1-25)")
	webFlag  = flag.Bool("web", false, "Start the web interface")

	allSolutions = make(map[string]map[string]func([]string) shared.Solution[any, any])
)

func init() {
	allSolutions["2022"] = y2022.Solutions
	allSolutions["2023"] = y2023.Solutions
	allSolutions["2024"] = y2024.Solutions
	allSolutions["2025"] = y2025.Solutions
}

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

	if solvers, ok := allSolutions[yearStr]; ok {
		if solver, ok := solvers[dayStr]; ok {
			if err := shared.Run(solver, os.Stdin, os.Stdout); err != nil {
				log.Fatalf("Error running solution: %v", err)
			}
		} else {
			log.Fatalf("Solution for day %s in year %s not found.", dayStr, yearStr)
		}
	} else {
		log.Fatalf("Solutions for year %s not found.", yearStr)
	}
}

func startWebServer() {
	fmt.Println("Starting web server on :8080...")
	mux := http.NewServeMux()

	// Static file handling
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

	// API routes
	mux.HandleFunc("GET /api/years", getYearsHandler)
	mux.HandleFunc("GET /api/years/{year}/days", getDaysHandler)
	mux.HandleFunc("POST /api/years/{year}/days/{day}", solveProblemHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getYearsHandler(w http.ResponseWriter, r *http.Request) {
	var years []string
	for year := range allSolutions {
		years = append(years, year)
	}
	sort.Strings(years)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(years)
}

func getDaysHandler(w http.ResponseWriter, r *http.Request) {
	year := r.PathValue("year")
	if days, ok := allSolutions[year]; ok {
		var dayKeys []string
		for day := range days {
			dayKeys = append(dayKeys, day)
		}
		sort.Strings(dayKeys)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dayKeys)
	} else {
		http.Error(w, fmt.Sprintf("Year %s not found or no solutions available.", year), http.StatusNotFound)
	}
}

func solveProblemHandler(w http.ResponseWriter, r *http.Request) {
	year := r.PathValue("year")
	day := r.PathValue("day")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	inputReader := strings.NewReader(string(body))

	solver, ok := allSolutions[year][day]
	if !ok {
		http.Error(w, fmt.Sprintf("Solution for day %s in year %s not found.", day, year), http.StatusNotFound)
		return
	}

	reader, writer := io.Pipe()
	defer reader.Close()

	go func() {
		defer writer.Close()
		if err := shared.Run(solver, inputReader, writer); err != nil {
			fmt.Fprintf(writer, "Error running solution: %v\n", err)
		}
	}()

	w.Header().Set("Content-Type", "text/plain")
	if _, err := io.Copy(w, reader); err != nil {
		log.Printf("Error streaming solution output to response: %v", err)
	}
}
