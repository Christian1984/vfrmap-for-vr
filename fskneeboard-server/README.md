# msfs2020-go

simconnect package [msfs2020-go/simconnect](simconnect/) connects to microsoft flight simulator 2020 using golang.

cross-compiles from macos/linux, no other dependencies required. produces a single binary with no other files or configuration required.

## how to build

- install sass compiler via npm
- install go version 1.14.14
- install https://github.com/jteeuwen/go-bindata globally
- install https://github.com/tc-hib/go-winres globally (go get https://github.com/tc-hib/go-winres)
- install https://github.com/boltdb/bolt globally (go get github.com/boltdb/bolt/...)
- copy required stuff to folder _vendor (see README.md there)
- run build-fskneeboard-server-FREE.bat or build-fskneeboard-server-PRO.bat

## status

[msfs2020-go/simconnect](simconnect/) package currently only implements enough of the simconnect api for [examples](examples/) and [vfrmap](vfrmap).

## releases and download

program zips releases are uploaded [here](https://github.com/lian/msfs2020-go/releases)

## tools

* [vfrmap](vfrmap/) local web-server that will allow you to view your location, and some information about your trajectory including airspeed and altitude.

## examples

* [examples/request_data](examples/request_data/) port of `MSFS-SDK/Samples/SimConnectSamples/RequestData/RequestData.cpp`

## Why does my virus-scanning software think this program is infected?

From official golang website https://golang.org/doc/faq#virus

"This is a common occurrence, especially on Windows machines, and is almost always a false positive. Commercial virus scanning programs are often confused by the structure of Go binaries, which they don't see as often as those compiled from other languages."
