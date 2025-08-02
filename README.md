# lsof_go

`lsof_go` is a simplified reimplementation of the Unix `lsof` command, written in [Go](https://golang.org). It lists currently running processes and their associated metadata (like PID, command name,username and commandline) by reading from the Linux `/proc` filesystem.

> ⚠️ **Note**: Works only on Linux. `/proc` is not available on macOS or Windows.

---

## 🚀 Features

- List active process IDs and command names
- Optionally show the username (-username)
- Lightweight and fast
- Educational — great for understanding how `/proc` works

---

## 📦 Installation

Clone the repo:

```bash
git clone https://github.com/YOUR_USERNAME/golsof.git
cd golsof
go build -o golsof
