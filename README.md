# FSKneeboard: A powerful VR-Cockpit Manager for Microsoft Flight Simulator

This project (formerly known as VFR Map For VR) adds a helpful ingame panel to your flight simulator which brings

- Several different Maps, Waypoint<sup>\*</sup> and Tracks<sup>\*</sup>,
- a Charts Viewer<sup>\*</sup> and
- a Notepad<sup>\*</sup>

right into your cockpit! This is is especially helpful for those of us who like to fly in VR (and thus not being able to look on a physical kneeboard, tablet, or second screen).

Additionally, FSKneeboard adds a very helpful fully automated and configurable Autosave Feature<sup>\*</sup>.

I made this mod for myself and for now, it does exactly what I want. And since the latest release, it does even more stuff that you guys, the community, asked for!

If you like it, please let me know and share it with other VR pilots :-) Also, consider yourself invited to join me and the community on the [FSKneeboard Discord Server](https://discord.fskneeboard.com) at https://discord.fskneeboard.com

*(\*) indicates "premium"-features that are available in FSKneeboard PRO, which is available for a Pay-What-You-Want-Price at https://fskneeboard.com/buy-now*

---

# Table of Contents

1. [Screenshots](#screenshots)
2. [TL;DR](#tldr)
3. [Support Your Modders :-)](#support-your-modders-smiley)
4. [Go PRO](#go-pro)
5. [Features](#features)
6. [Components](#components)
7. [Installation](#installation)
8. [Usage](#usage)
9. [Advanced Usage](#advanced-usage)
10. [Troubleshooting](#troubleshooting)
11. [Roadmap](#roadmap)
12. [Attribution](#attribution)
13. [Releases and Downloads](#releases-and-downloads)
14. [How to Contribute?](#how-to-contribute)
15. [HELP!!! Why Does My Virus-Scanning Software Think This Program Is Infected?](#help-why-does-my-virus-scanning-software-think-this-program-is-infected)

---

<div style="page-break-after: always;"></div>

# Screenshots

![Toolbar Icon](screenshots/fskneeboard-1.png)
![Navigational Data Enabled](screenshots/fskneeboard-2.jpg)
![Charts Viewer](screenshots/fskneeboard-3.jpg)
![Notepad](screenshots/fskneeboard-4.jpg)

# TL;DR

The Mod consists of **TWO PARTS(!)**: a *server* and an *ingame-panel* that you need to install and run **BOTH**!

**Please take the time to at least read the Installation and Usage sections below!!!**

It's dead simple! But if you only install the panel to the community folder and ignore the server this mod won't run and may appear "broken" to you!

**Some malware- and virus scanners detect FSKneeboard.exe as a virus! This is a false positive and known issue. Please read below ["HELP!!! Why Does My Virus-Scanning Software Think This Program Is Infected?"](#help-why-does-my-virus-scanning-software-think-this-program-is-infected) to learn more!**

---

# Support Your Modders :-)

If you enjoy this project, please consider buying me a coffee and/or donating to the guys I mentioned in the Attribution section. It allows us to keep developing addons and mods like these ones and making them available for free. Any amount is welcome! Thank you.

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate?hosted_button_id=ED8RR2JTV9BGU)

## Go PRO

Alternatively, you may decide to "go pro" at a Pay-What-You-Want-Price! It's really your call! You'll unlock all features of FSKneeboard plus you support the mod development. Also, you'll make a 1-year-old and a 5-year-old very, very happy, as I can buy them more ice cream :-)

---

<div style="page-break-after: always;"></div>

# Features

- The VFR-Map is a separate panel inside the sim: No fiddling around with virtual desktop browser windows etc.
- Map resolution etc. is optimized for VR use.
- Hide your own airplane on the map for a fully-fledged "paper map on kneeboard"-VFR navigation feeling
- Toggle to show and automatically follow your airplane on the map for a more "GPS"-ish style of navigation
- Five different map types
- Several navigation data overlays, including open flightmaps and openAIP
- Configurable a hotkey to show/hide the FSKneeboard panel while ingame.
- Add, remove and modify waypoints and tracks on the map<sup>\*</sup>
- Pull the currently loaded flightplan from MSFS into your kneeboard map, including ICAO identifiers<sup>\*</sup>
- Watch charts and checklists inside the integrated charts viewer<sup>\*</sup>
- Take notes inflight with your mouse on the integrated notepad<sup>\*</sup>
- Automatically create snapshots/savegames from your flights every few minutes so you're able to recover Flight Simulator instabilities and crashes (fully configurable, see [Advanced Usage Section](#advanced-usage) for details)<sup>\*</sup>

*(\*) indicates "premium"-features that are available in FSKneeboard PRO, which is available for a Pay-What-You-Want-Price at https://fskneeboard.com/buy-now*

---

# Components

The mod project consists of two components:

## Server

`fskneeboard-server` is the webserver that connects to your flight simulator and communicates with it to receive location data etc.

## Client

`fskneeboard-panel` is the actual ingame panel.

---

<div style="page-break-after: always;"></div>

# Installation

Download the zip from [here](https://github.com/Christian1984/vfrmap-for-vr/releases)

## Server

Place the contents of `fskneeboard-server` file somewhere convenient (like `C:\Tools\fskneeboard\`).

![Server Installation](screenshots/fskneeboard-install-server-1.png)

<div style="page-break-after: always;"></div>

Afterward, your `fskneeboard-server` folder should look like this:

![Server Installation - Done](screenshots/fskneeboard-install-server-2.png)

<div style="page-break-after: always;"></div>

If you have purchased FSKneeboard PRO make sure to also place your fskneeboard.lic-License file here! Your finished PRO-installation should look like this:

![Server Installation - FSKneeboard PRO - Done](screenshots/fskneeboard-pro-install-server-2.png)

<div style="page-break-after: always;"></div>

## Client

Place the folder `christian1984-ingamepanel-fskneeboard` in your MSFS community folder (typically `C:\Users\[username]\AppData\Local\Packages\Microsoft.FlightSimulator_8wekyb3d8bbwe\LocalCache\Packages\Community`)

![Panel Installation](screenshots/fskneeboard-install-panel-1.png)

<div style="page-break-after: always;"></div>

When finished, your Community folder should look like this (though there may be more than one extension installed, obviously).

![Panel Installation - Done](screenshots/fskneeboard-install-panel-2.png)

## Upgrading to a New Version

Please remove **all files** from the server directory and `christian1984-ingamepanel-fskneeboard` from your community folder, before installing a new version. Do not "copy over". No excuses! Old files may interfere with new ones and break the mod. Just do it, but keep your license file :-)

---

<div style="page-break-after: always;"></div>

# Usage

- Depending on where you have purchased Microsoft Flight Simulator, you may use launch FSKneeboard's starting:
    - `fskneeboard-autostart-windows-store.bat` for owners who have purchased via Windows Store (which simply calls `fskneeboard.exe --winstorefs`), or
    - `fskneeboard-autostart-steam.bat` for owners who have purchased via Steam (which simply calls `fskneeboard.exe --steamfs`)

![FSKneeboard - Autostart Scripts](screenshots/fskneeboard-autostart-shortcuts.png)

- If you encounter any errors or unexpected behaviour, simply run `fskneeboard.exe`. This has the autostart feature disabled by default. FSKneeboard will now wait until you have started Flight Simulator **manually**!
- Open up the ingame panel once inside the sim (like you would do for ATC etc.)
- Place conveniently in your VR space
- Click "Center Airplane" to initially center the map on your airplane.

<div style="page-break-after: always;"></div>

## Map Viewer

The map is the core component of FSKneeboard and available to FREE and PRO users alike. It contains several map modes as well as a representation of your own aircraft that you may also turn off so that you can "navigate by hand" on bush trips and the likes.

Owners of FSKneeboard PRO can also access the waypoint feature, which allows you to manually place waypoints on the map:

- A track will be automatically added between waypoints and a flag will show you information about the distance to the final waypoint and the heading of your track.
- You can click any given waypoint you have placed to toggle the visibility of its particular info-flag. The info-flag of the last waypoint of your track is always visible.
- If you get in the proximity of less than 0.5 NM of any given waypoint it will automatically be removed from the map. The last waypoint on your track will not be removed automatically, however, so you can use it to find a mission target or destination.

![Map Viewer](screenshots/fskneeboard-map-legend.png)

PRO users may also pull the currently loaded ingame-flightplan from their Flight Simulator onto the kneeboard by clicking the "cloud-icon" in the bottom left corner. This will load the flightplan you have created on the Worldmap screen before starting the flight. This will replace all manually placed waypoints on your map.

> PLEASE NOTE: When you change your flightplan by adding or removing waypoints through your ingame GPS, for example, these changes will not be reflected by the flightplan pulled from the Sim by this feature. For the time being, this is a known limitation of the feature. Please configure your flightplan before taking off on the Worldmap screen of MSFS.

<div style="page-break-after: always;"></div>

## Charts Viewer

FSKneeboard PRO contains a fully-fledged charts viewer for charts in png format. You can navigate the charts by either using the toolbar on the top or by dragging to pan the map around. You can also use your mouse wheel to zoom. 

![Charts Viewer](screenshots/fskneeboard-charts-legend.png)

<div style="page-break-after: always;"></div>

Make sure to place your charts inside the `charts` folder inside the server directory.

![Charts Folder](screenshots/fskneeboard-pro-charts-1.png)

<div style="page-break-after: always;"></div>

Your charts folder should look like this:

![Inside Your Charts Folder](screenshots/fskneeboard-pro-charts-2.png)

<div style="page-break-after: always;"></div>

## Notepad

Notepad is another feature that FSKneeboard PRO users have access to. It allows you to take notes during your flight session by simply drawing on it with your mouse.

![Map Viewer](screenshots/fskneeboard-notepad-legend.png)

## Autosave

Autosave is a feature that allows you to automatically create "snapshots" of your flights on predefined intervals. This is especially useful if you happen to encounter occasional (or even frequent) crashes to desktop (CTDs) with Microsoft Flight Simulator in VR.

Simply run FSKneeboard with the flag `--autosave [number]` to create a snapshot every `[number]` minutes. For example, run `fskneeboard.exe --autosave 5` to create one savegame every 5 minutes.

![Autosave Shortcut](screenshots/fskneeboard-autosave.png)

FSKneeboard automatically deletes older snapshots and keep only the latest 5.

If you need to restore a flight, you can find your autosaves inside your FSKneeboard-Server folder in the subdirectory `autosave`, e.g. `C:\Tools\fskneeboard\autosave`.

## Hotkey (Experimental)

> REMARKS:
> - This feature is an experimental feature. It's not bulletproof and may require some tweaking in the future. Use it at your own discretion and please provide feedback!
> - When the panel is hidden via the hotkey it is technically still there in your 3D space. Any attempt to control cockpit instrumentation that is hidden "behind" your invisible kneeboard will not work. The panel still intercepts your 3D cursor. Please make sure to position your kneeboard panel in a way that would allow you to interact with your cockpit instrumentation all the time, not matter if the panel is hidden or not!

You can define one of three hotkeys to toggle the visibility of the FSKneeboard ingame panel. The hotkey can be configured by starting FSKneeboard with the `--hotkey [number]` with `[number]` having the following meaning:

- `--hotkey 1` => `[Alt]+F`
- `--hotkey 2` => `[Alt]+K`
- `--hotkey 3` => `[Alt]+T`
- `--hotkey 4` => `[Ctrl]+[Shift]+F`
- `--hotkey 5` => `[Ctrl]+[Shift]+K`
- `--hotkey 6` => `[Ctrl]+[Shift]+T`

For example, launch `fskneeboard.exe --hotkey 1` to setup `[ALT]+F` as your FSKneeboard hotkey.

![Hotkey Shortcut](screenshots/fskneeboard-hotkey.png)

When ingame, you'll have to open the FSKneeboard panel ONCE by clicking it on the toolbar. For the rest of the flight, you can use the configured hotkey to toggle the panel's visibility as desired.

If you like, you can use your HOTAS configuration software to map this hotkey/shortcut to any button on your HOTAS.

---

<div style="page-break-after: always;"></div>

# Advanced Usage

The FSKneeboard server can be started with several commandline arguments to further customize its behaviour. In general, all you need to do is add them behind your "fskneeboard.exe" shortcut.

- `--autosave [number]`: Automatically create snapshots/savegames of your flights every `[number]` minutes.
- `--hotkey [number]`: Enable the hotkey feature and set it to Alt+F `[number] = 1`, Alt+K `[number] = 2` or Alt+T `[number] = 3`.
- `--winstorefs`: Start FSKneeboard together with your Flight Simulator purchased via Windows Store.
- `--steamfs`: Start FSKneeboard together with your Flight Simulator purchased via Steam.
- `--noupdatecheck`: Prevent FSKneeboard from checking the GitHub API for updates every three days.
- `--quietshutdown`: Prevent FSKneeboard from showing a "Press ENTER to continue..." prompt after disconnecting from MSFS

![Quietshutdown Shortcut](screenshots/fskneeboard-quietshutdown.png)

---

# Troubleshooting

- "I get errors when I try to start the server!" => This can happen if, for whatever reason, `fskneeboard.exe` cannot write `simconnect.dll`. Use your windows search to search for simconnect.dll (or download a copy somewhere on the interwebs) and copy it to the same directory `fskneeboard.exe` is located!
- "Windows says FSKneeboard contains a virus!" => That is a false positive and a well-known problem with software written in GO. Please make sure to read the section [HELP!!! Why Does My Virus-Scanning Software Think This Program Is Infected?](#help-why-does-my-virus-scanning-software-think-this-program-is-infected) below.
- "I've placed my pdf-charts in the charts directory but I can't see them inside the sim!" => Due to the limited capabilities of the browser engine that is embedded in Flight Simulator, the charts viewer can only display charts in png format. You will have to convert your charts. There is a multitude of pdf-to-png converters available online for free. Alternatively, you may want to take a look at GIMP, which is a freeware that also enables you to convert pdf files to png locally. If you know about other options, please reach out and let me know so that I can add them to this readme file.

---

# Roadmap

Here's a list of features that I've planned to implement in the foreseeable future: 

- Migrate the entire server component from GO to .NET to mitigate false virus alerts
- Integrate Navigraph and Little Nav Map (if feasible)
- Take notes via iPads and then have them synced to your in-sim notepad
- Allow for multiple notes to be taken, instead of having only one "sheet"
- Embed Twitch Chat / Discord / YouTube(?)
- (... This could be your wish :-) ...)

And here are some wishes from the community that I have to check for feasibility as soon as I get to it:

- Is it possible to add frequencies (such as VOR/DME/Comm)?
- Is it possible to change airport elevation data from [m] to [ft]?

---

# Attribution

This project uses forks of two amazing community projects. Without them, it would have taken me an incredible amount of time building this thing all on my own. Hence, I want to thank the two:

- The server is forked from [lian/msfs2020-go](https://github.com/lian/msfs2020-go).
- The client/ingame panel is forked from [bymaximus/msfs2020-toolbar-window-template](https://github.com/bymaximus/msfs2020-toolbar-window-template).

Great work, guys! Thanks for sharing your work with us!!!

Panel-Bar Icon made by [Freepik](https://www.freepik.com) from [www.flaticon.com](https://www.flaticon.com/).

Helicopter Icon provided by [SVGRepo](https://www.svgrepo.com/svg/128811/helicopter-bottom-view-silhouette).

Other Icons provided by [FontAwesome](https://fontawesome.com/license/free).

Maps and APIs:

- [OpenStreetMap](https://www.openstreetmap.org/copyright)
- [Stamen](http://maps.stamen.com)
- [Carto](https://carto.com/)
- [openAIP](https://www.openaip.net)
- [open flightmaps](https://www.openflightmaps.org/)

---

# Releases and Downloads

Program zips of FSKneeboard FREE are released and uploaded [here](https://github.com/Christian1984/vfrmap-for-vr/releases).

If you decide to support the development of this mod by buying a copy of FSKneeboard PRO [here](https://fskneeboard.com/buy-now) you will be emailed a link where you can download the FSKneeboard PRO binaries.

---

# How to Contribute?

If you have suggestions or issues, please feel free to reach out to me or create an issue within the Github repository. You may also add stuff yourself. Pull requests are very welcome!

---

# HELP!!! Why Does My Virus-Scanning Software Think This Program Is Infected?

From official golang website https://golang.org/doc/faq#virus

"This is a common occurrence, especially on Windows machines, and is almost always a false positive. Commercial virus scanning programs are often confused by the structure of Go binaries, which they don't see as often as those compiled from other languages."

Personal statement:

> If you don't trust the binary my suggestion would be two-fold:
>
> - Step 1: Upload the binaries to virustotal and see how many scanners throw a positive.
> - Step 2: Clone the repository and build the binary yourself. Everything is open-source, hence anyone who knows anything about building software can check the codebase for harmful code...
>
> Generally speaking (and that's true for anything you download from the web): If you don't trust the code, don't execute it! Especially NOT with elevated rights(!!!) I can understand anyone who doesn't want to run the software and appreciate how people ask questions instead of simply running the server...
>
> At the end of the day, it is your call. If you don't trust me and a binary that is flagged by a virus scanner, I do respect that. It's common sense, and I am not happy with the entire false-positive situation either. That is why in the mid-term I am planning to migrate the entire server component from GO to .NET, which should hopefully help to mitigate the issue.

However, here's a virus [report from virustotal for the binaries of Version 1.2.0](https://www.virustotal.com/gui/file/fc643b18493735e8e8e931c5434541738b650eca1692148bfaf57c04270f67b1/detection). Please note how many virus scanners do NOT flag the server as a positive. Please upload your own copy of FSKneeboard to virustotal to see the results of the latest version for yourself.

![Virustotal Report of Version 1.2.0](screenshots/virus-total-1-2-0.png)

---

# Support Your Modders :-)

If you enjoy this project, please consider buying me a coffee and/or donating to the guys I mentioned in the Attribution section above. It allows us to keep developing addons and mods like these ones and making them available for free. Any amount is welcome! Thank you.

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate?hosted_button_id=ED8RR2JTV9BGU)