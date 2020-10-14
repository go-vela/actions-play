package main

import "fmt"

var (
	// App represents the semantic version information for the application.
	App string
  // Arch represents the architecture information for the application.
  Arch string
	// Built represents the build date information for the application.
	Built string
	// Git represents the git information for the application.
	Git string
	// Go represents the go information for the application.
	Go string
)

func main() {
  fmt.Println("App: ", App)
  
  fmt.Println("Arch: ", Arch)

  fmt.Println("Built: ", Built)

  fmt.Println("Git: ", Git)

  fmt.Println("Go: ", Go)
}
