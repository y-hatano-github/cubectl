![cubectrl](screenshot.gif)
# 🧊 cubectrl
Don't mistype `kubectrl` as `cubectrl`...  
`cubectrl` renders a 3D cube in your terminal instead of controlling Kubernetes.

# ⚙️ Features
- 🧊 Renders a 3D cube instead of controlling Kubernetes.
- 🔄 Rotate the cube with arrow keys or `wasd`.
- 🔍 Zoom in/out with `z` and `x`.
- 🚫 Absolutely no Kubernetes functionality included.

# 💾 Download
Pre-built binaries are available for Windows, macOS, and Linux.

👉 Get the latest release here:
https://github.com/y-hatano-github/cubectrl/releases/latest

# 🚀 Quick start
## 🐧 Linux
```bash
wget https://github.com/y-hatano-github/cubectrl/releases/latest/download/cubectrl_linux_amd64.tar.gz
tar -xzvf cubectrl_linux_amd64.tar.gz
sudo mv cubectrl /usr/local/bin/
cubectrl
```
## 🍎 macOS
```bash
curl -LO https://github.com/y-hatano-github/cubectrl/releases/latest/download/cubectrl_darwin_amd64.tar.gz
tar -xzvf cubectrl_darwin_amd64.tar.gz
sudo mv cubectrl /usr/local/bin/
cubectrl
```
## 🪟 Windows
```
Invoke-WebRequest -OutFile cubectrl_windows_amd64.tar.gz https://github.com/y-hatano-github/cubectrl/releases/latest/download/cubectrl_windows_amd64.tar.gz
tar -xzvf cubectrl_windows_amd64.tar.gz
.\cubectrl.exe
```

# 📘 Usage
```
Usage: cubectrl [Flags]

Control cube in your terminal instead of controlling Kubernetes.

Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit

Flags:
  -h, --help    help for cubectrl
```