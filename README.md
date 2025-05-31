# SysAct

**SysAct** is a terminal‑user‑interface (TUI) application written in Go, built with the Bubble Tea framework and Lip Gloss styling. It provides a quick way to perform common system actions (logout, suspend, reboot, poweroff) directly from your terminal or launcher.

---

## 🚀 Features

* **Logout**: End the current user session and return to the login screen.
* **Suspend**: Pause the system, saving the session to memory and entering a low‑power state.
* **Reboot**: Restart the system, closing all applications and reinitializing the OS.
* **Poweroff**: Completely shut down the system.
* **Responsive UI**: Centrally‑aligned list menu that adapts to terminal window size.
* **Confirmation Dialog**: Ask for user confirmation (with countdown timer) before performing actions.
* **Localization**: Full support for English, French, Spanish, and Arabic; easily extendable via `config.toml`.
* **Custom Keybindings**: Remap navigation, confirm, cancel, and quit keys via `config.toml`.
* **Theme & Colors**: Configure `primary_color` and `secondary_color` in `config.toml` for Lip Gloss styles.
* **Structured Logging**: Built‑in multi‑level loggers with automatic rotation (via `lumberjack`).
* **Debian Package**: Pre‑built `.deb` available in GitHub Releases for easy installation.

---

## ⚙️ Installation

There are two main ways to install SysAct:

### 1. Install via Debian Package

A pre‑built `.deb` is provided in the [Releases](https://github.com/AyKrimino/SysAct/releases). Simply:

```sh
# Download the .deb for your architecture (e.g. sysact_0.1.0_amd64.deb)
sudo dpkg -i sysact_0.1.0_amd64.deb
```

Once installed, you can run:

```sh
sysact
```

> **Note:** This method is quickest for most Debian‑based distributions (Ubuntu, Debian, Mint, etc.).

### 2. Build from Source (Makefile)

Ensure you have Go 1.18+ installed and `$GOPATH/bin` (or Go’s bin directory) in your `PATH`. Then:

```sh
# Clone the repository
git clone https://github.com/AyKrimino/SysAct.git
cd SysAct

# Download dependencies
go mod download

# Build the executable into ./bin/sysact
make build

# You can now run:
./bin/sysact

# Optionally, add bin/ to your PATH:
export PATH="$PWD/bin:$PATH"
# So that typing "sysact" runs the app from anywhere:
sysact
```

---

## 🎮 Usage

After installation or building, simply execute:

```sh
sysact
```

* Use arrow keys (`↑`/`↓`) or `k`/`j` to navigate.
* Press `Enter` (or your configured `key_confirm`) to initiate an action.
* A confirmation dialog appears, with a countdown timer. Press `y` (or `key_confirm`) to confirm immediately, or wait for auto‑cancel.
* Press `q` or `Esc` (or your configured `key_quit`/`key_cancel`) to quit at any time.

---

## 📋 Configuration

SysAct reads settings from a TOML file located at:

```
~/.config/sysact/config.toml
```

If the file is missing, built‑in defaults are used. To customize, copy the example:

```sh
cp config.toml.example ~/.config/sysact/config.toml
```

Open `~/.config/sysact/config.toml` in your editor to adjust:

* **default\_language**: Choose among `en`, `fr`, `es`, `ar`.
* **confirm\_timeout**: Seconds before auto-confirm (1–60).
* **key\_up / key\_down / key\_confirm / key\_cancel / key\_quit / key\_select**: Customize navigation and control keys.
* **show\_help**: `true` or `false` to display or hide footer help text.
* **enable\_logging**: Enable file logging.
* **primary\_color / secondary\_color**: Hex codes (e.g. `#50fa7b`) for TUI styling.
* **\[lang.<code>]** sections: Localized UI strings for each supported language.

See `config.toml.example` for full annotated sample.

---

## 📂 Project Structure

```
├── LICENSE
├── Makefile                 # Build, test, and run shortcuts
├── README.md                # This file
├── bin/
│   └── sysact               # Pre‑built binary (ignored by Git)
├── cmd/
│   └── sysact/
│       └── main.go          # Entry point, loads config and starts TUI
├── config.toml.example      # Sample configuration with all keys commented
├── go.mod
├── go.sum
├── internal/
│   ├── config/
│   │   └── loader.go        # Loads + merges + validates config.toml
│   ├── logging/
│   │   └── logger.go        # Sets up multi‑level loggers with rotation
│   ├── system/
│   │   ├── actions.go       # Maps action indices to shell commands
│   │   └── utils.go         # Detects desktop environment / WM via XDG_CURRENT_DESKTOP
│   └── tui/
│       ├── action.go        # Defines a list item for a system action
│       ├── confirm_model.go # Manages confirm dialog state and timer
│       ├── confirm_update.go# Handles keypresses and timeout in confirm screen
│       ├── confirm_view.go  # Renders the confirm dialog
│       ├── main_model.go    # Main list model, holds items + config + styles
│       ├── main_update.go   # Updates main list (navigation + selection)
│       ├── main_view.go     # Renders the main screen
│       ├── root_model.go    # Switches between main and confirm models
│       └── styles.go        # Lip Gloss styles, now driven by config colors
└── logs/
    └── sysact.log           # Rotating log file (auto‑created & ignored by Git)
```

---

## 📦 Releases & Debian Package

Pre‑built Debian packages (`.deb`) are published on the [GitHub Releases](https://github.com/AyKrimino/SysAct/releases) page. To install:

```sh
# Download the latest sysact_<version>_amd64.deb
sudo dpkg -i sysact_<version>_amd64.deb
```

This installs `/usr/bin/sysact` so you can launch by typing `sysact` in any terminal.

---

## 🤝 Contributing

Contributions are welcome! To propose a change:

1. Fork the repository.
2. Create a new branch:  `git checkout -b feature/my-feature`
3. Make and commit your changes:  `git commit -m "feat: add my-feature"`
4. Push to your fork:  `git push origin feature/my-feature`
5. Open a pull request on the main repo.

Follow existing code style, add tests when appropriate, and update this README if you introduce new features.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/AyKrimino/SysAct/blob/main/LICENSE) file for details.
