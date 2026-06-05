# urlstash

A fast, minimal CLI for bookmarking URLs from your terminal. Save a URL under a short label, open it in your browser, copy it to your clipboard, or delete it — all without leaving the command line.

```
url add yt   https://youtube.com
url add gh   https://github.com
url open yt
url copy gh
url list
```

---

## Features

- Save URLs with short, memorable labels
- Open any saved URL directly in your default browser
- Copy a URL to the clipboard with a single command
- Delete entries you no longer need
- List all saved URLs in a clean, aligned table
- Colorized output for quick scanning
- Works on **Windows**, **macOS**, and **Linux** (X11 & Wayland)

---

## Installation

### Download from Releases (recommended)

Head to the [Releases page](https://github.com/RiseofRice/urlStash/releases) and download the binary for your platform:

| Platform | File to download |
|----------|-----------------|
| Windows (64-bit) | `urlstash_windows_amd64.exe` |
| macOS (Intel) | `urlstash_darwin_amd64` |
| macOS (Apple Silicon) | `urlstash_darwin_arm64` |
| Linux (64-bit) | `urlstash_linux_amd64` |
| Linux (ARM64) | `urlstash_linux_arm64` |

---

#### Windows

1. Download `url_windows_amd64.exe` and rename it to `url.exe`.
2. Move it to a folder on your `PATH`. A simple option is `C:\Windows\System32`:

   ```powershell
   # Run PowerShell as Administrator
   Move-Item .\url.exe C:\Windows\System32\url.exe
   ```

   **Or** add a custom folder to your `PATH`:

   ```powershell
   # Create a bin folder (once)
   New-Item -ItemType Directory -Force "$HOME\bin"

   # Move the binary there
   Move-Item .\url.exe "$HOME\bin\url.exe"

   # Add to PATH for current user (permanent)
   [Environment]::SetEnvironmentVariable(
     "PATH",
     "$env:PATH;$HOME\bin",
     "User"
   )
   ```

   Restart your terminal, then verify:

   ```powershell
   url version
   ```

---

#### macOS

1. Download the correct binary for your chip (`amd64` for Intel, `arm64` for Apple Silicon).
2. Rename it and make it executable:

   ```sh
   mv url_darwin_arm64 url          # or url_darwin_amd64
   chmod +x url
   ```

3. Move it onto your `PATH`:

   ```sh
   sudo mv url /usr/local/bin/url
   ```

   If you'd rather not use `sudo`, put it in `~/bin` instead:

   ```sh
   mkdir -p ~/bin
   mv url ~/bin/url

   # Bash — add to ~/.bashrc or ~/.bash_profile
   echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc && source ~/.bashrc

   # Zsh — add to ~/.zshrc
   echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc && source ~/.zshrc
   ```

4. On macOS you may need to allow the binary through Gatekeeper the first time:

   ```sh
   xattr -d com.apple.quarantine /usr/local/bin/url
   ```

   Or: right-click the file in Finder → **Open** → **Open** to approve it once.

5. Verify:

   ```sh
   url version
   ```

---

#### Linux

1. Download `url_linux_amd64` (or `url_linux_arm64` for ARM).
2. Rename, make executable, and move it onto your `PATH`:

   ```sh
   mv url_linux_amd64 url
   chmod +x url
   sudo mv url /usr/local/bin/url
   ```

   Without `sudo`, use a user-local bin directory:

   ```sh
   mkdir -p ~/.local/bin
   mv url ~/.local/bin/url

   # Add to PATH if not already present (Bash)
   echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc && source ~/.bashrc

   # Or for Zsh
   echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc && source ~/.zshrc
   ```

3. Verify:

   ```sh
   url version
   ```

---

### Build from source

You need [Go](https://go.dev/dl/) installed (1.21 or later).

```sh
git clone https://github.com/RiseofRice/urlStash.git
cd urlStash
go build -o url .
```

Move the binary somewhere on your `PATH` using the same steps shown above for your platform.

---

## Commands

### `add`

Save a URL under a label. Labels must be unique.

```
url add <label> <url>
```

```sh
url add docs    https://go.dev/doc
url add gh      https://github.com
url add yt      https://youtube.com
```

---

### `list` / `ls`

List all saved labels and their URLs in a formatted table.

```
url list
url ls
```

```
 LABEL  URL
 ──────────────────────────────────────────────────
 docs   https://go.dev/doc
 gh     https://github.com
 yt     https://youtube.com
```

---

### `open`

Open the URL for a label in your system's default browser.

```
url open <label>
```

```sh
url open yt
# Opens https://youtube.com in your browser
```

---

### `copy` / `cp` / `yank`

Copy the URL for a label to your clipboard.

```
url copy <label>
url cp   <label>
url yank <label>
```

```sh
url copy gh
# https://github.com is now in your clipboard
```

---

### `delete` / `del` / `rm`

Remove a saved label permanently.

```
url delete <label>
url del    <label>
url rm     <label>
```

```sh
url rm docs
```

---

### `version`

Print the current version.

```
url version
```

---

## Platform notes

### Clipboard dependencies

urlstash uses the native clipboard tool for each platform. On Linux you need one of the following installed:

| Desktop | Required tool | Install |
|---------|--------------|---------|
| X11     | `xclip`      | `sudo apt install xclip` / `sudo pacman -S xclip` |
| Wayland | `wl-copy`    | `sudo apt install wl-clipboard` / `sudo pacman -S wl-clipboard` |

macOS uses `pbcopy` (built in) and Windows uses `clip` (built in) — no extra dependencies needed.

### Opening URLs

| Platform | Tool used |
|----------|-----------|
| Linux    | `xdg-open` (install `xdg-utils` if missing) |
| macOS    | `open` (built in) |
| Windows  | `start` via `cmd` (built in) |

---

## Storage

URLs are stored as JSON in a platform-appropriate config directory:

| Platform | Path |
|----------|------|
| Windows  | `%APPDATA%\urlstash\store.json` |
| macOS    | `~/Library/Application Support/urlstash/store.json` |
| Linux    | `~/.config/urlstash/store.json` |

The file is plain JSON — you can inspect or back it up at any time.

---

## License

MIT — see [LICENSE](LICENSE).

---

## Acknowledgements

urlstash is built on the shoulders of two excellent Go libraries:

- [**spf13/cobra**](https://github.com/spf13/cobra) by Steve Francia

- [**fatih/color**](https://github.com/fatih/color) by Fatih Arslan
