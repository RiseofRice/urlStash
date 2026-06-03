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

### Build from source

You need [Go](https://go.dev/dl/) installed (1.21 or later).

```sh
git clone https://github.com/RiseofRice/urlStash.git
cd urlStash
go build -o url .
```

Move the binary somewhere on your `PATH`:

```sh
# Linux / macOS
mv url /usr/local/bin/url

# Windows (PowerShell — run as Administrator)
Move-Item url.exe C:\Windows\System32\url.exe
```

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
