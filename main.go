package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/aitsuki/avds/emulator"
)

var (
	Version = "dev"
)

func printVersion() {
	fmt.Printf(`avds %s
Go:      %s
OS/Arch: %s/%s
`,
		Version,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH)
}

func printHelp() {
	fmt.Print(`avds - Android Virtual Device Launcher

Usage:
  avds [flags]

Flags:
  -h, --help     Show help information
  -v, --version  Show version information

Examples:
  avds                    # Interactive AVD selection
  avds --version          # Show version info
`)
}

func main() {
	versionFlag := flag.Bool("version", false, "Show version information")
	shortVersionFlag := flag.Bool("v", false, "Show version information")
	helpFlag := flag.Bool("help", false, "Show help information")
	shortHelpFlag := flag.Bool("h", false, "Show help information")

	flag.Parse()

	if *versionFlag || *shortVersionFlag {
		printVersion()
		os.Exit(0)
	}

	// Handle help flags
	if *helpFlag || *shortHelpFlag {
		printHelp()
		os.Exit(0)
	}

	if err := emulator.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
