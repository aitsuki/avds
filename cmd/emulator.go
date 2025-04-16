package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetEmulatorPath() (string, error) {
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

func ListAvds(emulatorPath string) ([]string, error) {
	cmd := exec.Command(emulatorPath, "-list-avds")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	outputStr := strings.ReplaceAll(string(output), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(outputStr), "\n")
	var avds []string
	for _, line := range lines {
		if line != "" {
			avds = append(avds, line)
		}
	}

	if len(avds) == 0 {
		return nil, fmt.Errorf("no available avds")
	}
	return avds, nil
}

func StartAvd(emulatorPath, avdName string) error {
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
