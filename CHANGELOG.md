> REMARKS: changes indicated with (*) are exclusive features for "FSKneeboard PRO Supporters". If you want to support the development of the mod and unlock exclusive features, purchase a PRO license at https://fskneeboard.com/buy-now

# v1.3.0

## MAJOR:

- added access Flight Simulator's ingame flightplan from FSKneeboard and load it onto your kneeboard's map
- added a configurable hotkey to toggle the ingame panel's visibility from your keyboard and HOTAS (requires mapping of keyboard macros to your HOTAS buttons)
- added autoremoval for waypoints (except the last one) when getting within an 0.5 NM range

## MINOR:

- improved overview, performance and stability of the waypoints feature by disabling visibility of info-flags of individual waypoints (except the last one) by default. Individual info-flags can be toggled by clicking on the particular waypoint.
- added option `--quietshutdown` to prevent FSKneeboard from showing a "Press ENTER to continue..." prompt after disconnecting from MSFS
- added link to the FSKneeboard Discord server (https://discord.fskneeboard.com) to the server's startup message

---

# v1.2.1

## MINOR:

- hotfix for compatibility issues with other mods that add custom panels to MSFS

---

# v1.2.0

## MAJOR:

- added autosave feature to recover and continue your flights from system instabilities and Flight Simulator crashes (*)
- added session persistence to the charts viewer, i.e. the active chart and viewer state will be saved when the panel is closed and reopened during a flight (*)
- added dynamic resolution scaling and stretching
- added a helicopter icon for the chopper pilots out there
- added automatic update checker 

## MINOR:

- added separate autostart options for Flight Simulator versions purchased via Windows Store and Steam (`--winstorefs` and `--steamfs`)
- fixed map-behaviour when clicking on a track line while the teleport tool is selected
- fixed the appearance of the "Can't Load Map Message" that appears when the server has not been started
- improved the behaviour of the charts viewer toolbar when rescaling the FSKneeboard panel to a portrait-ish layout
- changed default visibility of NavData overlay to hidden
- load "freemium-info"-iframe content only when required
- MSFS SDK updated to v0.12.0

---

# v1.1.0

## MAJOR:

- "VFR Map For VR" was renamed to "FSKneeboard" to reflect the new scope of the extension
- Added ALL NEW Waypoints and Track, Charts Viewer, and Notepad for FSKneeboard PRO supporters
- FSKneeboard can now start MSFS in one go and then wait for MSFS to be ready. You do not have to tab out and start MSFS and FSKneeboard separately
- Entire project structure got a big overhaul under the hood

## MINOR:

- Added Font Awesome icons
- Reworked README
- Fixed NavData overlay persistence

---

# v1.0.3

## MAJOR:

- Fixed scaling issues introduced with MSFS world update 3
- Reworked "Center on Airplane" and "Show Airplane" UI elements

---

# v1.0.2

## MAJOR:

- show a warning message if the server isn't running

---

# v1.0.1

## MAJOR:

- first version of VFR Map For VR. minimum viable product