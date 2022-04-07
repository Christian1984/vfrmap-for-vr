package globals

import "vfrmap-for-vr/_vendor/premium/notepad"

var Pro bool

const DownloadLinkPro = "https://fskneeboard.com/download-latest"
const DownloadLinkFree = "https://fskneeboard.com/free-download/"

var DownloadLink string

var Quietshutdown bool
var DevMode bool
var ProductName string
var AutosaveInterval int
var HttpListen string
var LogLevel string

var SteamFs bool
var WinstoreFs bool
var MsfsAutostart bool
var ServerAutostart bool

var DisableTeleport = false

var DrmValid bool = false

var Verbose bool

var Notepad notepad.Notepad