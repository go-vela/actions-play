package main

import (
	"fmt"
	"runtime"
)

var (
	// Arch represents the architecture information for the application.
	Arch = runtime.GOARCH
	// BuildDate represents the build date information for the application.
	BuildDate string
	// GitBranch represents the git branch information for the application.
	GitBranch string
	// GitCommit represents the git commit information for the application.
	GitCommit string
	// GitState represents the git state information for the application.
	GitState string
	// GitSummary represents the git summary information for the application.
	GitSummary string
	// OS represents the operating system information for the application.
	OS = runtime.GOOS
	// Version represents the semantic version information for the application.
	Version string
)

func main() {
	fmt.Println("Arch: ", Arch)

	fmt.Println("Build Date: ", BuildDate)

	fmt.Println("Git Branch: ", GitBranch)

	fmt.Println("Git Commit: ", GitCommit)

	fmt.Println("Git State: ", GitState)

	fmt.Println("Git Summary: ", GitSummary)

	fmt.Println("OS: ", OS)

	fmt.Println("Version: ", Version)
}
