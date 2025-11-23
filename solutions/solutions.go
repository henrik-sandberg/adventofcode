package solutions

import (
	"adventofcode/solutions/shared"
	"adventofcode/solutions/y2022"
	"adventofcode/solutions/y2023"
	"adventofcode/solutions/y2024"
	"adventofcode/solutions/y2025"
	"sort"
)

var (
	internalSolutions = make(map[date]func([]string) shared.Solution[any, any])
)

type date struct {
	year string
	day  string
}

func init() {
	loadSolutions("2022", y2022.Solutions)
	loadSolutions("2023", y2023.Solutions)
	loadSolutions("2024", y2024.Solutions)
	loadSolutions("2025", y2025.Solutions)
}

func loadSolutions(year string, solutions map[string]func([]string) shared.Solution[any, any]) {
	for day, solver := range solutions {
		internalSolutions[date{year: year, day: day}] = solver
	}
}

// AvailableYears returns a slice of unique, sorted years.
func AvailableYears() []string {
	yearSet := make(map[string]any)
	for key := range internalSolutions {
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
	for key := range internalSolutions {
		if key.year == year {
			days = append(days, key.day)
		}
	}
	sort.Strings(days)
	return days
}

// GetSolution returns the solver function for a given year and day.
func GetSolution(year, day string) (func([]string) shared.Solution[any, any], bool) {
	solver, ok := internalSolutions[date{year: year, day: day}]
	return solver, ok
}
