# FSKneeboard Server

The FSKneeboard server connects to MSFS 2020 via the simconnect package [fskneeboard-server/simconnect](simconnect/) using Go.

## How to Build

- install go version 1.14.14

- install https://github.com/jteeuwen/go-bindata globally
- install https://github.com/tc-hib/go-winres globally (`go get https://github.com/tc-hib/go-winres`)
- install https://github.com/boltdb/bolt globally (`go get github.com/boltdb/bolt/...`)
- install https://github.com/shirou/gopsutil globally (`go get github.com/shirou/gopsutil`)
- install https://github.com/Christian1984/go-maptilecache globally (`go get github.com/Christian1984/go-maptilecache`)
- install https://github.com/Christian1984/go-update-checker globally (`go get github.com/Christian1984/go-update-checker`)

- install fyne.io/fyne/v2 globally (`go get fyne.io/fyne/v2`)
- if fyne dependencies won't build: 
  - download `x86_64-posix-seh` from https://sourceforge.net/projects/mingw-w64/files/ 
  - copy contents to `C:\MinGW-w64\`
  - add `C:\MinGW-w64\bin\` to PATH as described here https://dev.to/gamegods3/how-to-install-gcc-in-windows-10-the-easier-way-422j

- copy required stuff to folder _vendor (see README.md there)
- run build-fskneeboard-server-FREE.bat or build-fskneeboard-server-PRO.bat

## Why does my virus-scanning software think this program is infected?

From official golang website https://golang.org/doc/faq#virus

"This is a common occurrence, especially on Windows machines, and is almost always a false positive. Commercial virus scanning programs are often confused by the structure of Go binaries, which they don't see as often as those compiled from other languages."
