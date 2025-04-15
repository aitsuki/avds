# 获取最新版本
$latest = Invoke-RestMethod -Uri "https://api.github.com/repos/aitsuki/avds/releases/latest"
$version = $latest.tag_name

# 确定架构 (仅支持64位)
if (![Environment]::Is64BitOperatingSystem) {
    Write-Host "Error: only support amd64/arm64" -ForegroundColor Red
    exit 1
}

# 确定处理器架构 (amd64 或 arm64)
$arch = "amd64"
if ([System.Runtime.InteropServices.RuntimeInformation]::ProcessArchitecture -eq [System.Runtime.InteropServices.Architecture]::Arm64) {
    $arch = "arm64"
}

# 下载地址
$downloadUrl = "https://github.com/aitsuki/avds/releases/download/$version/avds-$version-windows-$arch.exe"

# 安装位置
$installDir = "$env:LOCALAPPDATA\avds"
$installPath = "$installDir\avds.exe"

# 创建安装目录
if (!(Test-Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir | Out-Null
}

# 下载并安装
Write-Host "Downloading avds $version (windows/$arch)..."
Invoke-WebRequest -Uri $downloadUrl -OutFile $installPath

# 添加到 PATH
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (!$userPath.Contains($installDir)) {
    [Environment]::SetEnvironmentVariable("Path", "$userPath;$installDir", "User")
    Write-Host "Added $installDir to PATH environment variable"
}

Write-Host "Installation completed! Please reopen your command prompt or PowerShell window to use the 'avds' command"