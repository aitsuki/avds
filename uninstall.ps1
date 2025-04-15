# Installation directory
$installDir = "$env:LOCALAPPDATA\avds"
$installPath = "$installDir\avds.exe"

# Check if installed
if (!(Test-Path $installPath)) {
    Write-Host "avds is not installed at $installPath"
    exit 0
}

# Remove binary and directory
Write-Host "Removing avds..."
Remove-Item -Path $installPath -Force
if ((Get-ChildItem -Path $installDir -Force | Measure-Object).Count -eq 0) {
    Remove-Item -Path $installDir -Force
    Write-Host "Removed empty directory $installDir"
}

# Check PATH
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath.Contains($installDir)) {
    $newPath = ($userPath -split ';' | Where-Object { $_ -ne $installDir }) -join ';'
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "Removed $installDir from PATH environment variable"
}

Write-Host "Uninstallation completed!"