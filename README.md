![cubectl](screenshot.gif)
# 🧊 cubectl
Don't mistype `kubectl` as `cubectl`...  
`cubectl` renders a 3D cube in your terminal instead of controlling Kubernetes.

# ⚙️ Features
- 🧊 Renders a 3D cube instead of controlling Kubernetes.
- 🔄 Rotate the cube with arrow keys or `wasd`.
- 🔍 Zoom in/out with `z` and `x`.
- 🚫 Absolutely no Kubernetes functionality included.

# 💾 Download
Pre-built binaries are available for Windows, macOS, and Linux.

👉 Get the latest release here:
https://github.com/y-hatano-github/cubectl/releases/latest

# 🚀 Quick start
## 🐧 Linux
```bash
wget https://github.com/y-hatano-github/cubectl/releases/latest/download/cubectl_linux_amd64.tar.gz
tar -xzvf cubectl_linux_amd64.tar.gz
sudo mv cubectl /usr/local/bin/
cubectl
```
## 🍎 macOS
```bash
curl -LO https://github.com/y-hatano-github/cubectl/releases/latest/download/cubectl_darwin_amd64.tar.gz
tar -xzvf cubectl_darwin_amd64.tar.gz
sudo mv cubectl /usr/local/bin/
cubectl
```
## 🪟 Windows
```
Invoke-WebRequest -OutFile cubectl_windows_amd64.tar.gz https://github.com/y-hatano-github/cubectl/releases/latest/download/cubectl_windows_amd64.tar.gz
tar -xzvf cubectl_windows_amd64.tar.gz
.\cubectl.exe
```

# 📘 Usage
```
Usage: cubectl [Flags]

Control cube in your terminal instead of controlling Kubernetes.

Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit

Flags:
  -h, --help    help for cubectl
```