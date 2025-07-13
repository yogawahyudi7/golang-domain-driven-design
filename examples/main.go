//go:build examples
// +build examples

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run -tags=examples examples/*.go <demo_name>")
		fmt.Println("Available demos:")
		fmt.Println("  - pretty_logging")
		fmt.Println("  - log_naming")
		fmt.Println("  - date_rotation")
		fmt.Println("  - log_rotation")
		fmt.Println("  - tomorrow_simulation")
		return
	}

	switch os.Args[1] {
	case "pretty_logging":
		runPrettyLoggingDemo()
	case "log_naming":
		runLogNamingDemo()
	case "date_rotation":
		runDateRotationDemo()
	case "log_rotation":
		runLogRotationDemo()
	case "tomorrow_simulation":
		runTomorrowSimulationDemo()
	default:
		fmt.Printf("Unknown demo: %s\n", os.Args[1])
	}
}
