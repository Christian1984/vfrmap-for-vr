package globals

import (
	"time"
	"vfrmap-for-vr/_vendor/premium/notepad"
)

var Pro bool

const DownloadLinkPro = "https://fskneeboard.com/download-latest"
const DownloadLinkFree = "https://fskneeboard.com/free-download/"

var DownloadLink string

var Quietshutdown bool
var DevMode bool

var ProductName string
var BuildVersion string

var AutosaveInterval int
var HttpListen string
var LogLevel string

var SteamFs bool
var WinstoreFs bool
var MsfsAutostart bool

const MaptileCacheMaxMemoryUsageDefault int = 8 * 128 * 1024 * 1024
const MaptileCacheTimeToLiveDefault time.Duration = 45 * 24 * time.Hour

var MaptileCacheMaxMemoryUsage int

var DisableTeleport = false

var DrmValid bool = false

var Verbose bool

var Notepad notepad.Notepad

var WipeMaptileCaches bool

var TourIndexStarted bool
var TourMapStarted bool
var TourChartsStarted bool
var TourNotepadStarted bool
var TourGuiStarted bool
