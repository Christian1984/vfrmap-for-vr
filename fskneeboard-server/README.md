# FSKneeboard Server

The simconnect package [fskneeboard-server/simconnect](simconnect/) connects to MSFS 2020 using Go.

## How to Build

- install sass compiler via npm
- install go version 1.14.14
- install https://github.com/jteeuwen/go-bindata globally
- install https://github.com/tc-hib/go-winres globally (go get https://github.com/tc-hib/go-winres)
- install https://github.com/boltdb/bolt globally (go get github.com/boltdb/bolt/...)
- copy required stuff to folder _vendor (see README.md there)
- run build-fskneeboard-server-FREE.bat or build-fskneeboard-server-PRO.bat

## Why does my virus-scanning software think this program is infected?

From official golang website https://golang.org/doc/faq#virus

"This is a common occurrence, especially on Windows machines, and is almost always a false positive. Commercial virus scanning programs are often confused by the structure of Go binaries, which they don't see as often as those compiled from other languages."
