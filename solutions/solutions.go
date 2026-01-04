package solutions

import (
	"adventofcode/solutions/shared"
	"adventofcode/solutions/y2021"
	"adventofcode/solutions/y2022"
	"adventofcode/solutions/y2023"
	"adventofcode/solutions/y2024"
	"adventofcode/solutions/y2025"
	"sort"
)

var (
	internalSolvers = make(map[date]shared.Solver)
)

type date struct {
	year string
	day  string
}

func init() {
	loadSolvers("2021", y2021.Solvers)
	loadSolvers("2022", y2022.Solvers)
	loadSolvers("2023", y2023.Solvers)
	loadSolvers("2024", y2024.Solvers)
	loadSolvers("2025", y2025.Solvers)
}

func loadSolvers(year string, solvers map[string]shared.Solver) {
	for day, solver := range solvers {
		internalSolvers[date{year: year, day: day}] = solver
	}
}

// AvailableYears returns a slice of unique, sorted years.
func AvailableYears() []string {
	yearSet := make(map[string]any)
	for key := range internalSolvers {
		yearSet[key.year] = nil
	}
	var years []string
	for year := range yearSet {
		years = append(years, year)
	}
	sort.Strings(years)
	return years
}

// AvailableDays returns a slice of unique, sorted days for a given year.
func AvailableDays(year string) []string {
	var days []string
	for key := range internalSolvers {
		if key.year == year {
			days = append(days, key.day)
		}
	}
	sort.Strings(days)
	return days
}

// GetSolver returns the solver function for a given year and day.
func GetSolver(year, day string) (shared.Solver, bool) {
	solver, ok := internalSolvers[date{year: year, day: day}]
	return solver, ok
}
