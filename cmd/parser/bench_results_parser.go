package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	Date   int64   `json:"date"`
	Suites []Suite `json:"suites"`
}

type Suite struct {
	Goos       string      `json:"goos"`
	Goarch     string      `json:"goarch"`
	Pkg        string      `json:"pkg"`
	Benchmarks []Benchmark `json:"benchmarks"`
}

type Benchmark struct {
	Name    string  `json:"Name"`
	Runs    int64   `json:"Runs"`
	NsPerOp float64 `json:"NsPerOp"`
	Mem     Mem     `json:"Mem"`
}

type Mem struct {
	AllocsPerOp float64 `json:"AllocsPerOp"`
	BytesPerOp  float64 `json:"BytesPerOp"`
	MBPerSec    float64 `json:"MBPerSec"`
}

func main() {
	bts, err := os.ReadFile("README.md")
	if err != nil {
		panic(err)
	}

	readme := string(bts)

	// clear readme after ## Challenges
	markdown := strings.Split(readme, "## Challenges")[0]

	files, err := os.ReadFile("bench.json")
	if err != nil {
		panic(err)
	}

	var results []Result
	if err := json.Unmarshal(files, &results); err != nil {
		panic(err)
	}

	markdown += "\n"
	markdown += "## Challenges\n"
	markdown += "\n"

	for _, result := range results {
		for _, suite := range result.Suites {
			markdown += fmt.Sprintf("* [%s](#%s) \n", getChallengeNameFromPackage(suite.Pkg), strings.ToLower(getChallengeNameFromPackage(suite.Pkg)))
		}
	}

	markdown += "\n"

	for _, result := range results {
		markdown += "\n"
		for _, suite := range result.Suites {
			markdown += fmt.Sprintf("\n## [%s](./%s)\n", getChallengeNameFromPackage(suite.Pkg), getSubPackageName(suite.Pkg))
			markdown += "\n"

			for i, benchmark := range suite.Benchmarks {
				if i == 0 {
					markdown += fmt.Sprintf("Runs: %d  \n", benchmark.Runs)
					markdown += "\n"
					markdown += "| Name | ns/op | B/op | allocs/op |  \n"
					markdown += "| ---- | ----- | ---- | --------- |  \n"
				}

				markdown += fmt.Sprintf("| %s | %s | %s | %s |  \n", benchmark.Name, fmt.Sprintf("%f ns/op", benchmark.NsPerOp), fmt.Sprintf("%f B/op", benchmark.Mem.BytesPerOp), fmt.Sprintf("%f allocs/op", benchmark.Mem.AllocsPerOp))
			}
		}
	}

	os.WriteFile("README.md", []byte(markdown), 0644)
}

func getSubPackageName(pk string) string {
	return strings.Replace(pk, "leet-code/", "", 1)
}

func getChallengeNameFromPackage(pkg string) string {
	return strings.ToTitle(strings.Replace(getSubPackageName(pkg), "-", " ", -1))
}
