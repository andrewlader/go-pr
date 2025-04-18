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

	switch stringState {
	case "open":
		state = goprlib.StateOpen
	case "closed":
		state = goprlib.StateClosed
	case "merged":
		state = goprlib.StateMerged
	case "all":
		state = goprlib.StateAll
	default:
		state = goprlib.StateOpen
	}
}

func handleExit() {
	recovery := recover()
	if recovery != nil {
		errOutput := fmt.Sprintf("panic occurred:\n    %v", recovery)
		goprlib.PrintError(errOutput)
		goprlib.PrintError("go-pr has stopped with an error")

		os.Exit(1)
	} else if finishedSuccessfully {
		var stringState string

		switch state {
		case goprlib.StateAll:
			stringState = "all of the"
		case goprlib.StateClosed:
			stringState = "all of the closed"
		case goprlib.StateMerged:
			stringState = "all of the merged"
		case goprlib.StateOpen:
			stringState = "all of the open"
		}

		goprlib.PrintError(fmt.Sprintf("go-pr has listed %s PRs for repo \"%s\" successfully", stringState, repo))
	}
}
