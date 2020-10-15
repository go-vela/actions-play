package main

import (
	"fmt"
	"runtime"
)

var (
	// App represents the semantic version information for the application.
	App string
	// Arch represents the architecture information for the application.
	Arch = runtime.GOARCH
	// Built represents the build date information for the application.
	Built string
	// Git represents the git information for the application.
	Git string
	// Lang represents the go information for the application.
	Lang string
	// OS represents the operating system information for the application.
	OS = runtime.GOOS
)

func main() {
	fmt.Println("App: ", App)

	fmt.Println("Arch: ", Arch)

	fmt.Println("Built: ", Built)

	fmt.Println("Git: ", Git)

	fmt.Println("Lang: ", Lang)

	fmt.Println("OS: ", OS)
}
