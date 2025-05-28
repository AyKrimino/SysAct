# SysAct

**SysAct** is a terminal‑user‑interface (TUI) application written in Go, built with the Bubble Tea framework and Lip Gloss styling. It provides a quick way to perform common system actions (logout, suspend, reboot, poweroff) directly from your terminal or launcher.

---

## 🚀 Features

* **Logout**: End the current user session and return to the login screen.
* **Suspend**: Pause the system, saving the session to memory and entering a low‑power state.
* **Reboot**: Restart the system, closing all applications and reinitializing the OS.
* **Poweroff**: Completely shut down the system.
* **Responsive UI**: Centrosymmetric list menu that adapts to terminal window size.
* **Structured Logging**: Built‑in log rotation and different log levels via `lumberjack`.

---

## 🔧 Installation

### Prerequisites

Go 1.18+ installed and $GOPATH/bin (or Go’s bin directory) in your PATH.

### Building from Source

Rather than including a compiled binary in the repository, users can easily build SysAct locally:

```bash
git clone https://github.com/AyKrimino/SysAct.git
cd SysAct

# Install or update dependencies
go mod download

# Build the executable into ./bin/
go build -o bin/sysact ./cmd/sysact

# Optionally add bin/ to your PATH for easy invocation
export PATH="$PWD/bin:$PATH"

# Now you can launch SysAct from anywhere:
sysact
```

---

## 🎮 Usage

Simply type:

```bash
sysact
```

Use the arrow keys (or `j`/`k`) to navigate and `Enter` (or space) to execute the selected action.

Press `q` or `Ctrl+C` to quit.

---

## ⚙️ Configuration

SysAct detects your desktop environment or window manager via the `XDG_CURRENT_DESKTOP` environment variable. If detection fails or you need custom commands, you can override defaults by creating a `config.toml` in the working directory (coming soon).

---

## 📂 Project Structure

```
LICENSE
README.md
cmd/
  └── sysact/
      └── main.go  Entry point
internal/
  ├── logging/
  │   └── logger.go   Rotating and leveled logging
  ├── system/
  │   ├── actions.go  Action initialization & execution
  │   └── utils.go    Desktop/WM detection
  └── tui/
      ├── action.go
      ├── model.go
      ├── styles.go
      ├── update.go
      └── view.go     
logs/
  └── sysact.log  Generated logs (ignored by git)
```

---

## 🤝 Contributing

1. Fork the repo.
2. Create a branch: `git checkout -b feature/my-awesome-feature`
3. Commit your changes: `git commit -m "feat: add my awesome feature"`
4. Push to your branch and open a pull request.

Please adhere to the existing code style and include tests where appropriate.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
