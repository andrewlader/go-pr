package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	_ "embed"

	"github.com/andrewlader/go-pr/internal/goprlib"
)

//go:embed git-describe.txt
var buildInfo string

var displayBuildInformation bool
var repo string
var state goprlib.StateFilter
var finishedSuccessfully bool

func init() {
	defer handleExit()

	finishedSuccessfully = false

	// parse any and all arguments provided
	parseArguments()
}

func main() {
	defer handleExit()

	if displayBuildInformation {
		buildInformation := strings.Split(buildInfo, "\n")
		if len(buildInformation) > 0 {
			goprlib.PrintVersionInfo("go-pr version: ", buildInformation[0])
		}
		if len(buildInformation) > 1 {
			goprlib.PrintVersionInfo("go version:      ", buildInformation[1])
		}
		if len(buildInformation) > 2 {
			goprlib.PrintVersionInfo("build date:      ", buildInformation[2])
		}
	} else {
		if len(repo) < 1 {
			panic("the repo flag is required; it defines which repo in the config to examine...")
		}
		finishedSuccessfully = true
	}
}

func parseArguments() {
	var stringState string

	flag.BoolVar(&displayBuildInformation, "version", false, "display build & version information")
	flag.StringVar(&repo, "repo", "", "defines the repo to examine (required)")
	flag.StringVar(&stringState, "state", "", "defines which PRs to list (optional, defaults to \"open\")")

	flag.Parse()

	state = goprlib.GetStateFromString(stringState)
}

func handleExit() {
	recovery := recover()
	if recovery != nil {
		errOutput := fmt.Sprintf("panic occurred:\n    %v", recovery)
		goprlib.PrintError(errOutput)
		goprlib.PrintError("go-pr has stopped with an error")

		os.Exit(1)
	} else if finishedSuccessfully {
		var output string

		stringState := state.ToString()
		if len(stringState) > 0 {
			output = fmt.Sprintf("go-pr has listed all of the %s PRs for repo \"%s\" successfully", stringState, repo)
		} else {
			output = fmt.Sprintf("go-pr has listed all of the PRs for repo \"%s\" successfully", repo)
		}

		goprlib.PrintError(output)
	}
}
