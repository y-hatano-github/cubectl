![cubectl demo: a rotating 3D cube in the terminal](screenshot.gif)
# 🧊 cubectl 
[![Release](https://img.shields.io/github/v/release/y-hatano-github/cubectl)](https://github.com/y-hatano-github/cubectl/releases/latest)
![Go Version](https://img.shields.io/github/go-mod/go-version/y-hatano-github/cubectl)
![License](https://img.shields.io/github/license/y-hatano-github/cubectl)

Once you mistype `kubectl` as `cubectl`, you’ll end up controlling a cube instead of Kubernetes clusters.  
This tool makes your typos fun.  

# ⚙️ Features
- 🧊 Renders a 3D cube instead of controlling Kubernetes clusters.
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
```powershell
Invoke-WebRequest -OutFile cubectl_windows_amd64.tar.gz https://github.com/y-hatano-github/cubectl/releases/latest/download/cubectl_windows_amd64.tar.gz
tar -xzvf cubectl_windows_amd64.tar.gz
.\cubectl.exe
```

# 📘 Usage
```
Usage: cubectl [Flags]

Control a cube in your terminal instead of controlling Kubernetes.

Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit

Flags:
  -h, --help    help for cubectl
```

# 📚 References
## Inspiration for the idea
I was reading the following articles, and somehow ended up creating this joke command.  
- https://www.reddit.com/r/kubernetes/comments/7tzyla/poll_how_to_pronounce_kubectl/

## For 3D rendering
- https://codezine.jp/article/detail/38?p=3
- https://qiita.com/mochimkchiking/items/24fbbe93e0f7aa89edba
