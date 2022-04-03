package globals

import "vfrmap-for-vr/freemium_src/notepad"

var Pro bool

const DownloadLinkPro = "https://fskneeboard.com/download-latest"
const DownloadLinkFree = "https://fskneeboard.com/free-download/"

var DownloadLink string

var Quietshutdown bool
var DisableTeleport bool
var DevMode bool
var ProductName string
var Hotkey int
var AutosaveInterval int
var HttpListen string

var SteamFs bool
var WinstoreFs bool

var DrmValid bool = false

var Verbose bool

var Notepad notepad.Notepad