# MSFS: VFR-MAP FOR VFR FLYING

This project adds a helpful VFR map as an ingame panel, which is especially helpful for those who like to fly in VR (and thus not being able to look on a physical kneeboard, tablet or second screen).

I made this mod for myself and for now, it does exactly what I want. If you like it. Let me know and share it with other VR pilots :-)

# Features

- VFR as a separate panel inside the sim: No fiddling around with virtual desktop browser windows etc.
- Map resolution etc. is optimized for VR use. Fair warning: On desktop browsers the map may look quite low res and UI elements may appear too big.
- Hide your own airplane on the map for a fully fledged "paper map on kneeboard"-VFR-navigation feeling
- Toggle to show and automatically follow your airplane on the map for a more "GPS"-ish style of navigation
- Five different map types
- Navigation data overlay

# Components

The mod projects consists of two components:

## Server

`vfrmap-server` is the webserver that connects to your flight simulator and communicates with it to receive location data etc.

## Client

`msfs-panel` is the actual ingame panel.

# How to install?

Download the zip from [here](https://github.com/Christian1984/vfrmap-for-vr/releases)

## Server

Place the `vfrmap.exe` file somewhere convenient (like C:\Tools\vfrmap\).

## Client

Place the folder `chris1984-ingamepanel-vfrmapforvr` in your MSFS community folder (typically `C:\Users\[username]\AppData\Local\Packages\Microsoft.FlightSimulator_8wekyb3d8bbwe\LocalCache\Packages\Community`)

# How to use?

- Start MSFS
- Start `vfrmap.exe`
- Open up the ingame panel once inside the sim (like you would do for ATC etc.)
- Place conveniently in your VR space

# Known Issues

- `vfrmap-server/build-vfrmap.sh` does not work properly at the moment. I somehow messed up the go module structure. For know, cd to `vfrmap-server/vfrmap` and run `go build` which should work, even though it does not generate the bindata.go files properly.

# Screenshots

// TODO

# Attribution

This project uses forks of two amazing community projects. Without them it would have taken me an incredible amount of time building this thing. Hence, I want to thank the two:

- The server is forked from [lian/msfs2020-go] (https://github.com/lian/msfs2020-go).
- The client/ingame panel is forked from [bymaximus/msfs2020-toolbar-window-template](https://github.com/bymaximus/msfs2020-toolbar-window-template).

Great work, guys! Thanks for sharing your work with us!!!

# Releases and Download

program zips releases are uploaded [here](https://github.com/Christian1984/vfrmap-for-vr/releases)

# How to contribute?

If you have suggestions or issues, please feel free to reach out to me or create an issue within the github repository. You may also add stuff yourself. Pull requests are very welcome!

# Why does my virus-scanning software think this program is infected?

From official golang website https://golang.org/doc/faq#virus

"This is a common occurrence, especially on Windows machines, and is almost always a false positive. Commercial virus scanning programs are often confused by the structure of Go binaries, which they don't see as often as those compiled from other languages."

# Support your modders :-)

If you enjoy this project, please consider buying me a coffee and/or donating to the guys I mentioned in the Attribution section above. It allows us to keep developing addons and mods like this ones and making them available for free. Any amount is welcome! Thank you.

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate?hosted_button_id=ED8RR2JTV9BGU)
