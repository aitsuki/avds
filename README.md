# avds - Android Virtual Device Launcher

Run Android emulator background from console.

A lightweight command-line tool for managing Android emulators with interactive selection and background launching.

## Requirements

- Android SDK installed
- `ANDROID_SDK_ROOT` or `ANDROID_HOME` environment variable set
- At least one AVD created via AVD Manager

## Usage

1. Run `avds` to list available devices
2. Select AVD using arrow keys
3. Press Enter to launch in background

## Development

```shell
git clone https://github.com/aitsuki/avds

# Build
cd avds
go build

# Run
go run .
```
