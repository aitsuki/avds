# avds - Android Virtual Device Launcher

[![GitHub release](https://img.shields.io/github/release/aitsuki/avds.svg)](https://github.com/aitsuki/avds/releases)
[![License](https://img.shields.io/github/license/aitsuki/avds.svg)](LICENSE)

Run Android emulator background from console.

A lightweight command-line tool for managing Android emulators with interactive selection and background launching.

## Highlights

- 🚀 Fast and lightweight command-line interface
- 🔍 Interactive device selection (keyboard nativation and filtering)
- 🏃‍♂️ Background emulator launching
- 🌈 Cross-platform support (Windows, macOS, Linux)
- 🔧 Simple installation and usage

## Requirements

- Android SDK installed
- `ANDROID_HOME` or `ANDROID_SDK_ROOT` environment variable set
- At least one AVD created via AVD Manager

## Installation

### Linux/macOS

```shell
curl -sfL https://raw.githubusercontent.com/aitsuki/avds/main/install.sh | sh
```

### Windows

```powershell
iwr -useb https://raw.githubusercontent.com/aitsuki/avds/main/install.ps1 | iex
```

### Using Go (all platforms)

```shell
go install github.com/aitsuki/avds@latest
```

### Manual Installation

Download the latest binary from [GitHub Releases](https://github.com/aitsuki/avds/releases).

## Usage

1. Run `avds` to list available devices
2. Select an AVD (keyboard nativation or filtering)
3. Press Enter to launch in background

> In Windows WSL, you can use `avds.exe` to run Android emulators from the
> Windows Android SDK, rather than using the Android SDK installed in WSL.

## Uninstallation

### Linux/macOS

```shell
curl -sfL https://raw.githubusercontent.com/aitsuki/avds/main/uninstall.sh | sh
```

### Windows

```powershell
iwr -useb https://raw.githubusercontent.com/aitsuki/avds/main/uninstall.ps1 | iex
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
