package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Version = "dev"
)

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

	emulatorPath, err := getEmulatorPath()
	if err != nil {
		log.Fatalln(err)
	}

	p := tea.NewProgram(initialModel(emulatorPath))
	m, err := p.Run()
	if err != nil {
		log.Fatalln(err)
	}

	finalModel := m.(model)
	if finalModel.selected != "" {
		err := startAvd(emulatorPath, finalModel.selected)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

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
