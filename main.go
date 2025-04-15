package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Version = "0.1.0"
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

type model struct {
	avds     []string
	cursor   int
	selected string
}

func initialModel(emulatorPath string) model {
	avds, err := listAvds(emulatorPath)
	if err != nil {
		log.Fatalln(err)
	}
	return model{avds: avds}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.avds)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.avds[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""
	for i, avd := range m.avds {
		cursor := " "
		if i == m.cursor {
			cursor = "→"
		}
		s += fmt.Sprintf("%s %s\n", cursor, avd)
	}
	return s
}

func getEmulatorPath() (string, error) {
	sdkRoot := os.Getenv("ANDROID_SDK_ROOT")
	if sdkRoot == "" {
		sdkRoot = os.Getenv("ANDROID_HOME")
	}
	if sdkRoot == "" {
		return "", fmt.Errorf("ANDROID_SDK_ROOT or ANDROID_HOME environment variable not set")
	}

	exeName := "emulator"
	if runtime.GOOS == "windows" {
		exeName += ".exe"
	}

	emulatorPath := filepath.Join(sdkRoot, "emulator", exeName)
	if _, err := os.Stat(emulatorPath); os.IsNotExist(err) {
		return "", fmt.Errorf("emulator not found at: %s", emulatorPath)
	}

	return emulatorPath, nil
}

func listAvds(emulatorPath string) ([]string, error) {
	cmd := exec.Command(emulatorPath, "-list-avds")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var avds []string
	for _, line := range lines {
		if line != "" {
			avds = append(avds, line)
		}
	}

	if len(avds) == 0 {
		return nil, fmt.Errorf("no avaliable avds")
	}
	return avds, nil
}

func startAvd(emulatorPath, avdName string) error {
	cmd := exec.Command(emulatorPath, "-avd", avdName)

	setProcessAttributes(cmd)

	nullFile, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	cmd.Stdout = nullFile
	cmd.Stderr = nullFile
	cmd.Stdin = nullFile

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start emulator: %v", err)
	}

	go cmd.Wait()
	return nil
}
