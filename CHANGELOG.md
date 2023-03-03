> REMARKS: changes indicated with (\*) are exclusive features for "FSKneeboard PRO Supporters". If you want to support the development of the mod and unlock exclusive features, purchase a PRO license at https://fskneeboard.com/buy-now

# v1.11.0

## MAJOR:

-   added Bing Maps (requires a custom, free API key to work, \*)
-   added configurable hotkeys for navigating to the maps, charts and notepad pages
-   added a heading indicator to the wind gauge
-   fixed a bug that broke the map module on external devices (e.g. tablet computers)

---

# v1.10.1

## HOTFIX:

-   fixed the issue of pilots being stuck in cologne

## MINOR:

-   potential fix for server startup issue if public IP address cannot be obtained by FSKneeboard
-   improved logging

---

# v1.10.0

## MAJOR:

-   added a trail line to the map module, can be toggled via GUI (\*)
-   added a distance and angle measuring tool (\*)
-   added custom API key section to the settings dialog to add custom API keys (replaces the legacy CLI argument)
-   added the option to disable the local maptile cache for services with custom API keys

## MINOR:

-   improved logging during startup to improve customer support

---

# v1.9.0

## MAJOR:

-   added a "product tour" to introduce new users to important core features of FSKneeboard
-   added favorites to the charts viewer so that users can open up to 5 charts and quickly navigate between them (\*)
-   added seamless charts rotation to charts viewer (\*)
-   added a PDF-to-PNG importer (\*)
-   improved charts viewer with zoom and rotation settings on a per-chart-basis (\*)
-   improved map tile cache performance by adding an in-memory caching mechanism and serving each cache on a separate host
-   fixed openAIP navigation map layer, enabled individual layers for different nav data (aiports, airspaces, navaids etc.)

## MINOR:

-   added "clickable" link to the "Server Ready at ..." label on the Control Panel
-   fixed the "FastLaunch" command line argument for the Windows Store MSFS autostarter
-   improved the layout of the charts viewer user interface
-   improved logging, added system stats to log output
-   improved installer so that it will no longer ask for a license if it is already in the intended install directory (\*)
-   improved repair FSKneeboard script by also clearing the maptile cache

---

# v1.8.0

## MAJOR:

-   added "North Up", "Track Up" and "Manual Map Rotation" options to the map module
-   added an "Disable Tools" option to the map module's mode selector
-   added scale to map
-   added a local tile cache for the map module to reduce network traffic

## MINOR:

-   added a prompt that enables the user to either keep or remove FSKneeboard user data (settings, license file, logs, cached data etc.) upon uninstall
-   disabled responsive layout ("horizontal menus")
-   grouped icons together into fly-out submenus to make menu more "compact"

---

# v1.7.0

## MAJOR:

-   FSKneeboard Core Application now ships with a Graphical User Interface (GUI)
-   improved the performance of the waypoint and track calculations, which directly improves the detection and performance of click and drag events on the map and overall system load
-   added opentopomap as an additional map provider

## MINOR:

-   added a new start menu item which allows users to reset the local FSKneeboard database to "factory settings"

---

# v1.6.1

## MAJOR:

-   hotfix for two bugs that prevented the waypoints and notepad module to not load previously saved waypoints and notes when certain parameters were not already initialized (\*)

---

# v1.6.0

## EXPERIMENTAL:

-   support for note-taking with physical keyboards (\*)

## MAJOR:

-   complete overhaul of the notepad functionality (\*)
-   added a new Type Mode to the notepad for typing notes with either a virtual onscreen keyboard or a physical keyboard (\*)
-   added the capability to take multiple notes (\*)
-   added persistence to notes, flightplans and waypoints, the active chart etc.: they are now saved on the local server and will be available across sessions (\*)
-   added persistence to configuration parameters such as the resolution scale and stretch etc.: They will now be saved on the local server and be available across sessions, so that they don't have to be reconfigured for every flight
-   added support for external devices (tablets, laptops, and local browser on second screen)
-   added support for notes taken via a tablet to be instantly synced to the FSKneeboard panel inside your VR cockpit

## MINOR:

-   added map-zoom buttons to the UI to better support VR controllers
-   added a confirmation dialog to the delete note button to prevent trashing notes by accident
-   added a logging mechanism that can create log files for better tech support and better debugging capabilities
-   fixed a bug that could make notes and waypoints disappear when the window was resized
-   fixed a bug that caused distances to be shown in miles, not nautical miles
-   minor bug fixes

---

# v1.5.0

## EXPERIMENTAL:

-   added a location finder to the map component that allows you to search the map for a street address or POI (\*)

## MAJOR:

-   complete redesign/overhaul of the user interface for a modern, much cleaner look and feel
-   added a document browser to the charts viewer that allows an improved organization of your individual charts in subfolders (\*)

## MINOR:

-   the "rubberband" from the airplane to your first waypoint can now be toggled by clicking your airplane (\*)
-   added color coded warnings to the installer to clearly indicate whether the community folder was properly identified or not
-   added dynamic resolution scaling to the documentation
-   added new screenshots to the documentation
-   minor bug fixes

---

# v1.4.1

## MAJOR:

-   fixed "missing MSFS toolbar icon"-bug that was introduced with Sim Update 5
-   changed ingame icon to reflect FSKneeboard's "brand identity"

## MINOR:

-   improved dark mode by setting blend mode to multiply

---

# v1.4.0

## EXPERIMENTAL:

-   added more hotkey options (Ctrl + Shift + F, Ctrl + Shift + K, Ctrl + Shift + T)

## MAJOR:

-   added an all new windows installer which installs the server to a dedicated directory of the user's choice and the ingame panel to the (automatically detected) community folder ("should work" (TM) for both the Steam and the Windows Store versions of MSFS)...
-   ... and also install your license file in the proper directory (\*)
-   added icao identifiers associated with waypoints that are pulled from the ingame flightplan (\*)
-   added open flightmaps integration as a separate nav data overlay
-   added dark mode and "red flashlight" mode

## MINOR:

-   removed magenta line from latest known aircraft position to the first waypoint when the aircraft is invisible (\*)
-   added capability to toggle the visibility of the final waypoint's nav info flag as well (\*)
-   fixed resolution scaling persistence
-   improved teleport UI

---

# v1.3.0

## EXPERIMENTAL:

-   added a configurable hotkey to toggle the ingame panel's visibility from your keyboard and HOTAS (requires mapping of keyboard macros to your HOTAS buttons)

## MAJOR:

-   added access Flight Simulator's ingame flightplan from FSKneeboard and load it onto your kneeboard's map (\*)
-   added autoremoval for waypoints (except the last one) when getting within an 0.5 NM range (\*)
-   added a wind direction and velocity indicator

## MINOR:

-   improved overview, performance and stability of the waypoints feature by disabling visibility of info-flags of individual waypoints (except the last one) by default. Individual info-flags can be toggled by clicking on the particular waypoint.
-   added option `--quietshutdown` to prevent FSKneeboard from showing a `Press ENTER to continue...` prompt after disconnecting from MSFS
-   added link to the FSKneeboard Discord server (https://discord.fskneeboard.com) to the server's startup message

---

# v1.2.1

## MINOR:

-   hotfix for compatibility issues with other mods that add custom panels to MSFS

---

# v1.2.0

## MAJOR:

-   added autosave feature to recover and continue your flights from system instabilities and Flight Simulator crashes (\*)
-   added session persistence to the charts viewer, i.e. the active chart and viewer state will be saved when the panel is closed and reopened during a flight (\*)
-   added dynamic resolution scaling and stretching
-   added a helicopter icon for the chopper pilots out there
-   added automatic update checker

## MINOR:

-   added separate autostart options for Flight Simulator versions purchased via Windows Store and Steam (`--winstorefs` and `--steamfs`)
-   fixed map-behaviour when clicking on a track line while the teleport tool is selected
-   fixed the appearance of the "Can't Load Map Message" that appears when the server has not been started
-   improved the behaviour of the charts viewer toolbar when rescaling the FSKneeboard panel to a portrait-ish layout
-   changed default visibility of NavData overlay to hidden
-   load "freemium-info"-iframe content only when required
-   MSFS SDK updated to v0.12.0

---

# v1.1.0

## MAJOR:

-   "VFR Map For VR" was renamed to "FSKneeboard" to reflect the new scope of the extension
-   Added ALL NEW Waypoints and Track, Charts Viewer, and Notepad for FSKneeboard PRO supporters
-   FSKneeboard can now start MSFS in one go and then wait for MSFS to be ready. You do not have to tab out and start MSFS and FSKneeboard separately
-   Entire project structure got a big overhaul under the hood

## MINOR:

-   Added Font Awesome icons
-   Reworked README
-   Fixed NavData overlay persistence

---

# v1.0.3

## MAJOR:

-   Fixed scaling issues introduced with MSFS world update 3
-   Reworked "Center on Airplane" and "Show Airplane" UI elements

---

# v1.0.2

## MAJOR:

-   show a warning message if the server isn't running

---

# v1.0.1

## MAJOR:

-   first version of VFR Map For VR. minimum viable product
